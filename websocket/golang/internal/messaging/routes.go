package messaging

import (
	"github.com/gin-gonic/gin"
)

type MessagingRoute struct {
	controller *MessagingController
}

func NewMessagingRoute(controller *MessagingController) *MessagingRoute {
	return &MessagingRoute{
		controller: controller,
	}
}

func (route *MessagingRoute) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", route.controller.HandleMessage)
}
