package delivery

import (
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type demoDelivery struct {
	demoHandler *handler.DemoHandler
}

func NewDemoDelivery(demoHandler *handler.DemoHandler) *demoDelivery {
	return &demoDelivery{
		demoHandler: demoHandler,
	}
}

func (d *demoDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	demoGroup := apiGroup.Group("demo")
	{
		demoGroup.GET("query", d.demoHandler.Query)
		demoGroup.DELETE("param/:id", d.demoHandler.Param)
		demoGroup.POST("body", d.demoHandler.Body)
		demoGroup.PATCH("header", d.demoHandler.Header)
	}
}
