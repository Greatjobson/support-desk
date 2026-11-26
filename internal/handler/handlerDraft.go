package handler

import (
	"net/http"

	"github.com/Greatjobson/support-desk/internal/model/message"
	"github.com/Greatjobson/support-desk/internal/model/ticket"
	"github.com/Greatjobson/support-desk/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessageHandler struct {
	service service.MessageService
}

func NewMessageHandler(service service.MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (s *MessageHandler) CreatMessage(c *gin.Context) {
	var msg message.Message
	if err := c.ShouldBind(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg.ID = uuid.NewString() //TODO мне интересно а стоит ли менять айдишку иммено тут или не стоит? а как же время сообщение? ну ладно время установить мы потом напишем

	isFirstMessage, err := s.service.NewMessage(msg.TiketID, msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if isFirstMessage {
		c.JSON(http.StatusCreated, gin.H{"message: pls wait message from agent": msg.ID})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": msg})
}

func (s *MessageHandler) GetMessages(c *gin.Context) {
	id := c.Param("id")

	data, err := s.service.LoadMessage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}

type TicketHandler struct {
	service service.TicketService
}

func NewTicketHandler(service service.TicketService) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

func (s *TicketHandler) NewTicket(c *gin.Context) {
	var t ticket.Ticket

	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.service.NewTicket(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ticket": t.ID})
}

func (s *TicketHandler) GetTicket(c *gin.Context) {
	id := c.Param("id")
	data, err := s.service.GetTicket(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errohgyguyuyr": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ticket": data})
}
