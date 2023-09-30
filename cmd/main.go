package main

import (
	handle "websocket/delivery"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	handle.Router(engine)

	engine.Run(":8080")
}
