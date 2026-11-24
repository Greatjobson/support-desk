package main

import (
	"github.com/Greatjobson/support-desk/internal/handler"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, messageHandler *handler.MessageHandler, ticketHandler *handler.TicketHandler) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/message", messageHandler.CreatMessage)
}
