package messaging

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type MessagingController struct {
	upgrader     *websocket.Upgrader
	clientManger *ClientManager
}

func NewMessagingController(upgrader *websocket.Upgrader, clientManager *ClientManager) *MessagingController {
	return &MessagingController{
		upgrader:     upgrader,
		clientManger: clientManager,
	}
}

func (mc *MessagingController) HandleMessage(c *gin.Context) {
	conn, err := mc.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := NewClient(conn)
	mc.clientManger.AddClient(client)
	go client.Read()
	go client.Write()

}
