package delivery

import (
	"go-backend/internal/common/cache"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/handler"
	"time"

	"github.com/gin-gonic/gin"
)

type articleDelivery struct {
	articleHandler *handler.ArticleHandler
	cache          *cache.Cache
	authMiddleware *middlewares.AuthMiddleware
}

func NewArticleDelivery(articleHandler *handler.ArticleHandler, cache *cache.Cache, authMiddleware *middlewares.AuthMiddleware) *articleDelivery {
	return &articleDelivery{
		articleHandler: articleHandler,
		cache:          cache,
		authMiddleware: authMiddleware,
	}
}

func (d *articleDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	articleGroup := apiGroup.Group("article")
	{
		articleGroup.POST("", d.articleHandler.Create)
		articleGroup.GET("", d.cache.CacheMemory(5*time.Second), d.articleHandler.FindAll)
		articleGroup.DELETE(":id", d.articleHandler.Delete)
	}
}
