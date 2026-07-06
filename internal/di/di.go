package dependency

import (
	"go-backend/ent"
	"go-backend/internal/common/cache"
	"go-backend/internal/common/env"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/common/rabbitmq"
	"go-backend/internal/common/socket"
	storage_impl "go-backend/internal/common/storage/impl"
	"go-backend/internal/delivery"
	"go-backend/internal/handler"
	"go-backend/internal/repository/repository_impl"
	"go-backend/internal/usecase/usecase_impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Injection(ginEngine *gin.Engine, entClient *ent.Client, gormClient *gorm.DB, env *env.Env, allowOrigins []string, rabbitmq *rabbitmq.RabbitMQ) {
	tokenUsecase := usecase_impl.NewTokenUsecase(env)

	cache := cache.NewCache(env)

	// repository
	articleRepository := repository_impl.NewArticleRepository(entClient, gormClient)
	userRepository := repository_impl.NewUserRepository(entClient)
	chatGroupRepository := repository_impl.NewChatGroupRepository(entClient)
	chatGroupMemberRepository := repository_impl.NewChatGroupMemberRepository(entClient)
	unitOfWorkRepository := repository_impl.NewUnitOfWorkRepository(entClient)
	chatMessageRepository := repository_impl.NewChatMessageRepository(entClient)

	localFileStorage := storage_impl.NewLocalFileStorage("public")
	cloudinaryFileStorage := storage_impl.NewCloudinaryStorage(env)

	// Socket
	chatUsecase := usecase_impl.NewChatUsecase(tokenUsecase, userRepository, chatGroupRepository, chatGroupMemberRepository, chatMessageRepository, unitOfWorkRepository)
	chatHandler := handler.NewChatHandler(chatUsecase)
	socket := socket.NewSocket(chatHandler)
	socket.Start(ginEngine, allowOrigins)

	authMiddleware := middlewares.NewAuthMiddleware(tokenUsecase, userRepository)

	articleUsecase := usecase_impl.NewArticleUsecase(articleRepository)
	articleHandler := handler.NewArticleHandler(articleUsecase)
	articleDelivery := delivery.NewArticleDelivery(articleHandler, cache, authMiddleware)

	demoUsecase := usecase_impl.NewDemoUsecase()
	demoHandler := handler.NewDemoHandler(demoUsecase)
	demoDelivery := delivery.NewDemoDelivery(demoHandler)

	authUsecase := usecase_impl.NewAuthUsecase(userRepository, tokenUsecase, env)
	authHandler := handler.NewAuthHandler(authUsecase, env)
	authDelivery := delivery.NewAuthDelivery(authHandler, authMiddleware)

	userUsecase := usecase_impl.NewUserUsecase(userRepository, localFileStorage, cloudinaryFileStorage)
	userHandler := handler.NewUserHandler(userUsecase)
	userDelivery := delivery.NewUserDelivery(userHandler, authMiddleware)

	chatGroupUsecase := usecase_impl.NewChatGroupUsecase(chatGroupRepository)
	chatGroupHandler := handler.NewChatGroupHandler(chatGroupUsecase)
	chatGroupDelivery := delivery.NewChatGroupDelivery(chatGroupHandler)

	chatMessageUsecase := usecase_impl.NewChatMessageUsecase(chatMessageRepository)
	chatMessageHandler := handler.NewChatMessageHandler(chatMessageUsecase)
	chatMessageDelivery := delivery.NewChatMessageDelivery(chatMessageHandler)

	orderUsecase := usecase_impl.NewOrderUsecase(rabbitmq)
	orderHandler := handler.NewOrderHandler(orderUsecase)
	orderDelivery := delivery.NewOrderDelivery(orderHandler, authMiddleware)

	rootDelivery := delivery.NewRootDelivery(demoDelivery, articleDelivery, authDelivery, userDelivery, chatGroupDelivery, chatMessageDelivery, orderDelivery)
	rootDelivery.RegisterRouter(ginEngine)
}
