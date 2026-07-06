package handler

import (
	"errors"
	"fmt"
	"go-backend/internal/common/constant"
	"go-backend/internal/common/env"
	"go-backend/internal/common/helpers"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	env         *env.Env
}

func NewAuthHandler(authUsecase usecase.AuthUsecase, env *env.Env) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
		env:         env,
	}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var body dto.AuthRegisterReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		if err == io.EOF {
			ctx.Error(response.NewBadRequestException("Body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	result, err := a.authUsecase.Register(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var body dto.AuthLoginReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		if err == io.EOF {
			ctx.Error(response.NewBadRequestException("Body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	result, err := a.authUsecase.Login(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}

	setTokenCookie(ctx, result.AccessToken, result.RefreshToken)

	response.Success(true, "", 0, ctx)
}

func (a *AuthHandler) GetInfo(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	fmt.Println("helpers.GetUser", user)

	result, err := a.authUsecase.GetInfo(ctx.Request.Context(), user)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(
		map[string]any{
			"user": result,
		}, "", 0, ctx)
}

func (a *AuthHandler) RefreshToken(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("accessToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			return
		}
		ctx.Error(response.NewBadRequestException())
		return
	}
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			return
		}
		ctx.Error(response.NewBadRequestException())
		return
	}

	result, err := a.authUsecase.RefreshToken(ctx.Request.Context(), accessToken, refreshToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Trường hợp 1: trả cả cặp token (accessToken, refreshToken)
	// kỹ thuận rotate
	// nếu người dùng không sử dụng trong vòng 24h, thì refreshToken không được làm mới => logout

	// Trường hợp 2: trả cả accessToken
	// từ khi người dùng login -> 24h sau -> logout

	setTokenCookie(ctx, result.AccessToken, result.RefreshToken)

	response.Success(result, "", 0, ctx)
}

func (a *AuthHandler) GoogleLogin(ctx *gin.Context) {

	result, err := a.authUsecase.GoogleLogin(ctx.Request.Context())
	if err != nil {
		ctx.Redirect(http.StatusFound, a.env.DomainFe+"/login?error="+err.Error())
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		constant.GOOGLE_OAUTH_STATE,
		result.State,
		60,
		"/",
		"",
		false,
		true,
	)
	redirect := ctx.Query("redirect")
	if redirect != "" {
		ctx.SetCookie(
			constant.GOOGLE_OAUTH_REDIRECT,
			redirect,
			60,
			"/",
			"",
			false,
			true,
		)
	}

	ctx.Redirect(http.StatusFound, result.Url)
}

func (a *AuthHandler) GoogleCallback(ctx *gin.Context) {
	stateGoogle := ctx.Query("state")
	if stateGoogle == "" {
		errMess := "state empty"
		fmt.Println(errMess)
		ctx.Redirect(http.StatusFound, a.env.DomainFe+"/login?error="+errMess)
		return
	}

	stateCookie, err := ctx.Cookie(constant.GOOGLE_OAUTH_STATE)
	if err != nil {
		errMess := err.Error()
		fmt.Println(errMess)
		ctx.Redirect(http.StatusFound, a.env.DomainFe+"/login?error="+errMess)
		return
	}

	code := ctx.Query("code")
	if code == "" {
		errMess := "code empty"
		fmt.Println(errMess)
		ctx.Redirect(http.StatusFound, a.env.DomainFe+"/login?error="+errMess)
		return
	}

	input := dto.AuthGoogleCallbackInput{
		StateGoogle: stateCookie,
		StateCookie: stateGoogle,
		Code:        code,
	}

	result, err := a.authUsecase.GoogleCallback(ctx.Request.Context(), input)
	if err != nil {
		errMess := err.Error()
		fmt.Println(errMess)
		ctx.Redirect(http.StatusFound, a.env.DomainFe+"/login?error="+errMess)
		return
	}

	setTokenCookie(ctx, result.AccessToken, result.RefreshToken)

	redirectUrl, _ := ctx.Cookie(constant.GOOGLE_OAUTH_REDIRECT)
	if redirectUrl != "" {
		ctx.Redirect(http.StatusFound, redirectUrl)
		return
	}

	ctx.Redirect(http.StatusFound, a.env.DomainFe)
}

func setTokenCookie(ctx *gin.Context, accessToken string, refreshToken string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		"accessToken",
		accessToken,
		0,
		"/",
		"",
		false,
		true,
	)
	ctx.SetCookie(
		"refreshToken",
		refreshToken,
		0,
		"/",
		"",
		false,
		true,
	)
}
