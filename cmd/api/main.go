package main

import (
	in_memory "github.com/Greatjobson/support-desk/internal/database/in-memory"
	"github.com/Greatjobson/support-desk/internal/handler"
	"github.com/Greatjobson/support-desk/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ticketRepository := in_memory.NewTicketIMStorage()
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	messageRepository := in_memory.NewMessageIMStorage()
	messageService := service.NewMessageService(messageRepository, ticketService)
	messageHandler := handler.NewMessageHandler(messageService)

	setupRouter(server, messageHandler, ticketHandler)
	server.Run(":8080")
}
