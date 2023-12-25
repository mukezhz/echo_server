package main

import (
	"github.com/mukezhz/echo_server/internal/messaging"
	"github.com/mukezhz/echo_server/pkg/infrastructure"
)

func main() {
	router := infrastructure.NewRouter()
	messaging.InitializeMessagingModule(router)
	router.Run(":8080")
}
