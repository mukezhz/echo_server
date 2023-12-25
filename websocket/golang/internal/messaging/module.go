package messaging

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/echo_server/pkg/infrastructure"
)

func InitializeMessagingModule(router *gin.Engine) {
	upgrader := infrastructure.NewUpgrader()
	newClientManger := NewClientManger()
	controller := NewMessagingController(upgrader, newClientManger)
	route := NewMessagingRoute(controller)

	wsRouter := router.Group("/ws")
	route.RegisterRoutes(wsRouter)
}
