package cache

import (
	"context"
	"fmt"
	"go-backend/internal/common/env"
	"go-backend/internal/common/helpers"
	"log"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	memoryStore *persist.MemoryStore
	redisStore  *persist.RedisStore
}

func NewCache(env *env.Env) *Cache {
	// dữ liệu cache dính chặt với BE
	// deploy restart BE là dữ liệu cache bị mất
	memoryStore := persist.NewMemoryStore(1 * time.Minute)

	// database phụ, dữ liệu cache muốn share cho các BE khác
	// BE restart thì không bị mất dữ liệu cache
	// docker run --name some-redis -d -p 6379:6379  redis redis-server --requirepass "12345"
	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     env.RedisAddr,
		Password: env.RedisPass,
	}))

	statusCmd := redisStore.RedisClient.Ping(context.Background())
	err := statusCmd.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ [REDIS] Connection To Redis Succcessfully", statusCmd.String())

	return &Cache{
		memoryStore: memoryStore,
		redisStore:  redisStore,
	}
}

func (c *Cache) CacheMemory(defaultExpire time.Duration) gin.HandlerFunc {
	return cache.CacheByRequestURI(
		c.memoryStore,
		defaultExpire,
		cache.WithCacheStrategyByRequest(customKey),
	)
}

func (c *Cache) CacheReđis(defaultExpire time.Duration) gin.HandlerFunc {
	return cache.CacheByRequestURI(
		c.redisStore,
		defaultExpire,
		cache.WithCacheStrategyByRequest(customKey),
	)

}

func customKey(ctx *gin.Context) (bool, cache.Strategy) {
	user, err := helpers.GetUser(ctx)

	key := fmt.Sprintf(
		"%s:%s",
		ctx.Request.Method,
		ctx.Request.URL.RequestURI(),
	)

	if err == nil && user != nil {
		key = fmt.Sprintf(
			"%d:%s:%s",
			user.ID,
			ctx.Request.Method,
			ctx.Request.URL.RequestURI(),
		)
	}

	return true, cache.Strategy{
		CacheKey: key,
	}
}
