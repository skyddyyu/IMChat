// routes/routes.go
package routes

import (
	"ImChat/src/controllers"
	"ImChat/src/db"
	"ImChat/src/handlers"
	"ImChat/src/middlewares"
	"ImChat/src/models"
	"ImChat/src/repositories"
	"ImChat/src/services"
	"ImChat/src/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRepo := repositories.NewUserRepository(db.DB)
		userService := services.NewUserService(userRepo)
		userController := controllers.NewUserController(userService)
		userRoutes.POST("/register", userController.RegisterUser)
		userRoutes.POST("/login", userController.LoginUser)
		userRoutes.GET("/list", middlewares.Auth(), userController.GetUserList)
		userRoutes.POST("/logout", middlewares.Auth(), userController.Logout)
	}

	chatRoomRoutes := router.Group("/chatroom", middlewares.Auth())
	{
		chatRoomRepo := repositories.NewChatRoomReposotory(db.DB)
		chatRoomService := services.NewChatRoomService(chatRoomRepo)
		chatRoomController := controllers.NewChatRoomController(chatRoomService)
		chatRoomRoutes.POST("/create", chatRoomController.CreateChatRoom)

		userRoomChatRepo := repositories.NewUserRoomChatRepository(db.DB)
		userRoomChatService := services.NewUserChatRoomService(userRoomChatRepo)
		userRoomChatController := controllers.NewUserRoomChatController(userRoomChatService)
		chatRoomRoutes.POST("/join", userRoomChatController.JoinChatRoom)
	}

	wsRoutes := router.Group("/ws", middlewares.Auth())
	{
		wsRoutes.GET("", func(c *gin.Context) {
			// 创建WebSocket连接
			webSocketInstance, err := ws.UpgradeWebSocketConnection(c)
			if err != nil {
				// 处理连接失败
				handlers.ServerError(c, err.Error())
				return
			}
			defer webSocketInstance.Close()
			go ws.CheckHeartbeat(webSocketInstance)

			// 处理用户信息和添加到连接映射
			if err := ws.HandleUserInfoAndAddToConnection(webSocketInstance, c); err != nil {
				// 处理用户信息和添加到连接映射失败
				handlers.ServerError(c, err.Error())
				return
			}
			for {
				// 处理接收到的消息
				messageType, p, err := webSocketInstance.ReadMessage()
				if err != nil {
					delete(models.Connection, webSocketInstance)
					handlers.ServerError(c, err.Error())
					continue
				}
				if messageType == websocket.TextMessage {
					// 收到消息，开启一个线程去执行
					go ws.HandleReceivedMessage(p, c)
				}
			}
		})
	}
}
