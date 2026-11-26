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

	api := r.Group("/api/message/v1")
	{
		api.GET("/:id", messageHandler.GetMessages)
		api.POST("/", messageHandler.CreatMessage)
	}
	api2 := r.Group("/api/v1/ticket")
	{
		api2.GET("/:id", ticketHandler.GetTicket)
		api2.POST("/users", ticketHandler.NewTicket)
	}
}
