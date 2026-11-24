package handler

import (
	"github.com/Greatjobson/support-desk/internal/model/message"
	"github.com/Greatjobson/support-desk/internal/model/ticket"
	"github.com/Greatjobson/support-desk/internal/service"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service service.MessageService
}

func NewMessageHandler(service service.MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (s *MessageHandler) CreatMessage(c *gin.Context) error {
	err := s.service.NewMessage(ticketID, message)
	if err != nil {
		return err
	}
	return nil
}
func (s *MessageHandler) GetMessage(ticketID string) ([]message.Message, error) {
	return s.service.LoadMessage(ticketID)
}

type TicketHandler struct {
	service service.TicketService
}

func NewTicketHandler(service service.TicketService) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

func (s *TicketHandler) NewTicket(ticket ticket.Ticket) error {
	return s.service.NewTicket(ticket)
}

func (s *TicketHandler) GetTicket(ticketID string) (ticket.Ticket, error) {
	return s.service.GetTicket(ticketID)
}
