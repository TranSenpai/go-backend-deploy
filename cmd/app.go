package main

import (
	"fmt"
	"go-backend/ent"
	"go-backend/internal/common/ent_client"
	"go-backend/internal/common/env"
	"go-backend/internal/common/gorm_client"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/common/rabbitmq"
	"go-backend/internal/common/response"
	"go-backend/internal/common/swagger"
	dependency "go-backend/internal/di"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	ginEngine *gin.Engine
	env       *env.Env
	entClient *ent.Client
	rabbitmq  *rabbitmq.RabbitMQ
}

func NewApp() *App {
	env := env.New()

	// Create a Gin router with default middleware (logger and recovery)
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(middlewares.ErrorHandler)
	ginEngine.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.Error(response.NewInternalServerErrorException())
		ctx.Abort()
	}))

	// relativePath: bí dánh dùng ở web, thay thế cho tên folder
	// root: folder thật ở may/server sẽ được truy cập khi dùng bí danh
	// ginEngine.Static("go-backend", "") // lộ code
	// ginEngine.Static("docx", "./public/docx") // example
	// ginEngine.Static("xlxx", "./public/xlxx") // example
	ginEngine.Static("images", "./public/images")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://google.com"}
	ginEngine.Use(cors.New(corsConfig))

	// ginEngine.Use(func(ctx *gin.Context) {
	// 	ctx.Header("access-control-allow-origin", "http://localhost:3000")
	// 	if ctx.Request.Method == "OPTIONS" {
	// 		ctx.Header("access-control-allow-headers", "Origin,Content-Length,Content-Type")
	// 		ctx.Header("access-control-allow-methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
	// 		ctx.Header("access-control-max-age", "43201")
	// 		ctx.AbortWithStatus(204)
	// 		return
	// 	}
	// })

	swagger.Start(ginEngine, env)

	// ginEngine.Use(middlewares.A)
	// ginEngine.Use(middlewares.B)
	// ginEngine.Use(middlewares.C)
	entClient := ent_client.New(env)
	gormClient := gorm_client.New(env)
	rabbitmq := rabbitmq.NewRabbitMQ(env)
	dependency.Injection(ginEngine, entClient, gormClient, env, corsConfig.AllowOrigins, rabbitmq)

	return &App{
		ginEngine: ginEngine,
		env:       env,
		entClient: entClient,
		rabbitmq:  rabbitmq,
	}
}

func (a *App) Start() {
	addr := fmt.Sprintf("%s:%s", a.env.Host, a.env.Port)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	a.ginEngine.Run(addr)
}
