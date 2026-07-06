package usecase_impl

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"go-backend/ent"
	"go-backend/internal/common/env"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

type authUsecase struct {
	userRepository repository.UserRepository
	tokenUsecase   usecase.TokenUsecase
	oauth2Config   *oauth2.Config
}

func NewAuthUsecase(userRepository repository.UserRepository, tokenUsecase usecase.TokenUsecase, env *env.Env) usecase.AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
		tokenUsecase:   tokenUsecase,
		oauth2Config: &oauth2.Config{
			ClientID:     env.GoogleClientId,
			ClientSecret: env.GoogleClientSecret,
			RedirectURL:  env.GoogleRedirectUrl,
			Scopes: []string{
				"openid",
				"email",
				"profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

// Register implements [usecase.AuthUsecase].
func (a *authUsecase) Register(ctx context.Context, body dto.AuthRegisterReq) (any, error) {
	// Kiểm tra user tồn tại hay chưa
	isExits, err := a.userRepository.ExitsByEmail(ctx, body.Email)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	// Nếu đã tồn tại thì trả lỗi
	if isExits {
		return nil, response.NewBadRequestException("Email đã có vui lòng đăng nhập")
	}

	// Mã hoá password
	// hash: băm (password) không dịch ngược
	// encryption: mã hoá chuyển dữ liệu
	hashPassByte, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	body.Password = string(hashPassByte)

	// tạo người dừng mới
	userNew, err := a.userRepository.CreateUser(ctx, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return userNew, nil
}

// Login implements [usecase.AuthUsecase].
func (a *authUsecase) Login(ctx context.Context, body dto.AuthLoginReq) (*dto.AuthLoginReturn, error) {
	user, err := a.userRepository.FindUserByEmail(ctx, body.Email)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	if user.Password == nil {
		return nil, response.NewBadRequestException("Vui lòng đăng nhập bằng google để cập nhật password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(body.Password))
	if err != nil {
		return nil, response.NewBadRequestException("Mật khẩu không chính xác")
	}

	// trả về accessToken và refreshToken
	accessToken, err := a.tokenUsecase.CreateAccessToken(user.ID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	refreshToken, err := a.tokenUsecase.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return &dto.AuthLoginReturn{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GetInfo implements [usecase.AuthUsecase].
func (a *authUsecase) GetInfo(ctx context.Context, user *ent.Users) (*ent.Users, error) {
	return user, nil
}

// RefreshToken implements [usecase.AuthUsecase].
func (a *authUsecase) RefreshToken(ctx context.Context, accessToken string, refreshToken string) (*dto.AuthRefreshTokenReturn, error) {
	// jwt.WithoutClaimsValidation(): không kiểm tra hết hạn
	claimAccessToken, err := a.tokenUsecase.VerifyAccessToken(accessToken, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	claimRefreshToken, err := a.tokenUsecase.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	if claimAccessToken.UserId != claimRefreshToken.UserId {
		return nil, response.NewUnauthorizedException("2 Token không cùng 1 user")
	}

	user, err := a.userRepository.FindUserById(ctx, claimAccessToken.UserId)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	accessTokenNew, err := a.tokenUsecase.CreateAccessToken(user.ID)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	refreshTokenNew, err := a.tokenUsecase.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, response.NewUnauthorizedException(err.Error())
	}

	return &dto.AuthRefreshTokenReturn{
		AccessToken:  accessTokenNew,
		RefreshToken: refreshTokenNew,
	}, nil
}

// GoogleLogin implements [usecase.AuthUsecase].
func (a *authUsecase) GoogleLogin(ctx context.Context) (*dto.AuthGoogleLoginReturn, error) {
	bytes := make([]byte, 10)
	fmt.Println(bytes)

	rand.Read(bytes)
	fmt.Println(bytes)

	state := hex.EncodeToString(bytes)
	fmt.Println(state)

	url := a.oauth2Config.AuthCodeURL(state)
	fmt.Println("url", url)

	return &dto.AuthGoogleLoginReturn{
		State: state,
		Url:   url,
	}, nil
}

// GoogleCallback implements [usecase.AuthUsecase].
func (a *authUsecase) GoogleCallback(ctx context.Context, input dto.AuthGoogleCallbackInput) (*dto.AuthLoginReturn, error) {

	fmt.Println(input.Code)
	fmt.Println(input.StateCookie)
	fmt.Println(input.StateGoogle)

	if input.StateCookie != input.StateGoogle {
		return nil, errors.New("state invalid")
	}

	token, err := a.oauth2Config.Exchange(ctx, input.Code)
	if err != nil {
		return nil, err
	}

	id_token := token.Extra("id_token").(string)

	idtokenPayload, err := idtoken.Validate(ctx, id_token, a.oauth2Config.ClientID)
	if err != nil {
		return nil, err
	}

	email := idtokenPayload.Claims["email"].(string)
	email_verified := idtokenPayload.Claims["email_verified"].(bool)
	name := idtokenPayload.Claims["name"].(string)
	picture := idtokenPayload.Claims["picture"].(string)
	googleId := idtokenPayload.Claims["sub"].(string)

	// fmt.Println("token", token)
	// fmt.Printf("type = %T value %v", id_token, id_token)
	// fmt.Printf("\n\n type = %T value %+v \n\n", idtokenPayload, idtokenPayload)
	fmt.Println("email", email)
	fmt.Println("email_verified", email_verified)
	fmt.Println("name", name)
	fmt.Println("picture", picture)
	fmt.Println("googleId", googleId)

	if !email_verified {
		return nil, errors.New("email not verify")
	}

	userExists, err := a.userRepository.FindUserByEmail(ctx, email)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if userExists == nil {
		input := dto.AuthCreateUserForGoogleReq{
			Email:    email,
			FullName: name,
			Avatar:   picture,
			GoogleId: googleId,
		}
		userExists, err = a.userRepository.CreateUserForGoogle(ctx, input)
		if err != nil {
			return nil, err
		}
	}

	accessToken, err := a.tokenUsecase.CreateAccessToken(userExists.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.tokenUsecase.CreateRefreshToken(userExists.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthLoginReturn{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
