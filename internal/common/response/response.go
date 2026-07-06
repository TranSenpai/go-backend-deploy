package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessFormat[T any] struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
	Doc        string `json:"doc"`
}

func Success(data any, message string, code int, ctx *gin.Context) {
	if code == 0 {
		code = http.StatusOK
	}
	if message == "" {
		message = http.StatusText(code)
	}

	reuslt := SuccessFormat[any]{
		Status:     "success",
		StatusCode: code,
		Message:    message,
		Data:       data,
		Doc:        "",
	}
	ctx.JSON(reuslt.StatusCode, reuslt)
}

type ErrorFormat struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Doc        string `json:"doc"`
}

func Error(message string, code int, ctx *gin.Context) {
	ctx.JSON(code, ErrorFormat{
		Status:     "error",
		StatusCode: code,
		Message:    message,
		Doc:        "",
	})
}
