package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tinkerbaj/chatwebsocketgin/chat"
	"log"
)

func main() {

	//create new Hub and run it
	hub := chat.NewHub()
	go hub.Run()

	//we need pass hub to out route with roomid
	app := gin.Default()
	app.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		chat.ServeWS(c, roomId, hub)

	})

	log.Println(app.Run(":8080"))
}
