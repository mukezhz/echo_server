package messaging

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	Egress chan string
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn:   conn,
		Egress: make(chan string, 1),
	}
}

func (c *Client) Read() {
	defer c.Conn.Close()
	c.Conn.SetReadLimit(512)
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("ERROR:", err)
			break
		}
		c.Egress <- string(message)
		fmt.Println(string(message))
	}
	fmt.Println("EXITING READ")
}

func (c *Client) Write() {
	defer c.Conn.Close()
	fmt.Println("STARTING WRITE", c.Egress)
	for msg := range c.Egress {
		if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			fmt.Println("ERROR WRITING MESSAGE:", err)
			return
		}
		fmt.Println("SENT MESSAGE:", string(msg))
	}
}

// ClientManager is a struct that manages clients
type ClientManager struct {
	sync.RWMutex
	Clients []*Client
}

func NewClientManger() *ClientManager {
	return &ClientManager{
		Clients: make([]*Client, 0),
	}
}

func (cm *ClientManager) AddClient(client *Client) {
	cm.Lock()
	defer cm.Unlock()
	cm.Clients = append(cm.Clients, client)
}

func (cm *ClientManager) RemoveClient(conn *websocket.Conn) {
	cm.Lock()
	defer cm.Unlock()
	for i, client := range cm.Clients {
		if client.Conn == conn {
			cm.Clients = append(cm.Clients[:i], cm.Clients[i+1:]...)
			break
		}
	}
}
