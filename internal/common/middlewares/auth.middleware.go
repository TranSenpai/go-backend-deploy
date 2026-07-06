package middlewares

import (
	"errors"
	"fmt"
	"go-backend/internal/common/response"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	tokenUsecase   usecase.TokenUsecase
	userRepository repository.UserRepository
}

func NewAuthMiddleware(tokenUsecase usecase.TokenUsecase, userRepository repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		tokenUsecase:   tokenUsecase,
		userRepository: userRepository,
	}
}

func (a *AuthMiddleware) Protect(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("accessToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.Error(response.NewUnauthorizedException())
			ctx.Abort()
			return
		}
		ctx.Error(response.NewBadRequestException())
		ctx.Abort()
		return
	}

	claim, err := a.tokenUsecase.VerifyAccessToken(accessToken)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, jwt.ErrTokenExpired) {
			ctx.Error(response.NewForbiddenException())
			ctx.Abort()
			return
		}
		ctx.Error(response.NewUnauthorizedException())
		ctx.Abort()
		return
	}

	user, err := a.userRepository.FindUserById(ctx, claim.UserId)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		ctx.Abort()
		return
	}

	ctx.Set("user", user)

}
