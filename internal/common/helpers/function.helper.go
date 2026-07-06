package helpers

import (
	"errors"
	"go-backend/ent"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) (*ent.Users, error) {
	userAny, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("API chưa gắn middleware protect")
	}
	if userAny == nil {
		return nil, errors.New("Không có user")
	}
	user, ok := userAny.(*ent.Users)
	if !ok {
		return nil, errors.New("Type user không hợp lệ")
	}
	return user, nil
}
