package delivery

import (
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type searchDelivery struct {
	searchHandler *handler.SearchHandler
}

func NewSearchDelivery(searchHandler *handler.SearchHandler) *searchDelivery {
	return &searchDelivery{
		searchHandler: searchHandler,
	}
}

func (d *searchDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	searchGroup := apiGroup.Group("search")
	{
		searchGroup.GET("", d.searchHandler.FindAll)
	}
}
