package usecase_impl

import (
	"errors"
	"go-backend/internal/common/env"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenUsecase struct {
	env *env.Env
}

func NewTokenUsecase(env *env.Env) usecase.TokenUsecase {
	return &tokenUsecase{
		env: env,
	}
}

// CreateAccessToken implements [usecase.TokenUsecase].
func (t *tokenUsecase) CreateAccessToken(userId int) (string, error) {
	claim := dto.CustomClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.env.ExpiresAtAccessToken)),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return tokenClaim.SignedString([]byte(t.env.SecretAccessToken))
}

// VerifyAccessToken implements [usecase.TokenUsecase].
func (t *tokenUsecase) VerifyAccessToken(accessToken string, options ...jwt.ParserOption) (*dto.CustomClaim, error) {
	claim := &dto.CustomClaim{}
	tokenClaim, err := jwt.ParseWithClaims(
		accessToken,
		claim,
		func(*jwt.Token) (any, error) {
			return []byte(t.env.SecretAccessToken), nil
		},
		options...,
	)
	if err != nil {
		return nil, err
	}
	if !tokenClaim.Valid {
		return nil, errors.New("Access Token InValid không hợp lệ")
	}
	return claim, nil
}

// CreateRefreshToken implements [usecase.TokenUsecase].
func (t *tokenUsecase) CreateRefreshToken(userId int) (string, error) {
	claim := dto.CustomClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.env.ExpiresAtRefreshToken)),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return tokenClaim.SignedString([]byte(t.env.SecretRefreshToken))
}

// VerifyRefreshToken implements [usecase.TokenUsecase].
func (t *tokenUsecase) VerifyRefreshToken(refreshToken string, options ...jwt.ParserOption) (*dto.CustomClaim, error) {
	claim := &dto.CustomClaim{}
	tokenClaim, err := jwt.ParseWithClaims(
		refreshToken,
		claim,
		func(*jwt.Token) (any, error) {
			return []byte(t.env.SecretRefreshToken), nil
		},
		options...,
	)
	if err != nil {
		return nil, err
	}
	if !tokenClaim.Valid {
		return nil, errors.New("Refresh Token InValid không hợp lệ")
	}
	return claim, nil
}
