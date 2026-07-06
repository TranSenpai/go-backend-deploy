package middlewares

import (
	"errors"
	"go-backend/internal/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler captures errors and returns a consistent JSON error response
func ErrorHandler(c *gin.Context) {
	c.Next() // Process the request first

	// Check if any errors were added to the context
	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err

		statusCode := http.StatusInternalServerError
		message := http.StatusText(statusCode)

		// kiểm tra xem có phải là lỗi exception (lỗi custom)
		var exception *response.Exception
		isException := errors.As(err, &exception)
		if isException {
			statusCode = exception.StatusCode
			message = exception.Message
		}

		response.Error(message, statusCode, c)
	}
}
