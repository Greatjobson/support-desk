package service

import (
	"errors"
	"fmt"

	in_memory "github.com/Greatjobson/support-desk/internal/database/in-memory"
	"github.com/Greatjobson/support-desk/internal/model/message"
	"github.com/Greatjobson/support-desk/internal/model/ticket"
)

type TicketService interface {
	GetTicket(ticketID string) (ticket.Ticket, error)
	NewTicket(t ticket.Ticket) error
}

type ticketService struct {
	repo in_memory.TicketRepository
}

func NewTicketService(repo in_memory.TicketRepository) TicketService {
	return &ticketService{repo: repo}
}

func (s *ticketService) GetTicket(ticketID string) (ticket.Ticket, error) {
	t, err := s.repo.LoadTicket(ticketID)
	if err != nil {
		return ticket.Ticket{}, errors.New("error with load ticket") // TODO надо какой то стандарт придумать чтоб весь проект юзал один способ ерроры писать
	}
	return t, nil
}
func (s *ticketService) NewTicket(t ticket.Ticket) error {
	return s.repo.StoreTicket(t.ID, t)
}

type MessageService interface {
	NewMessage(ticketID string, mess message.Message) (bool, error)
	LoadMessage(ticketID string) ([]message.Message, error)
}

type messageService struct {
	repo in_memory.MessageRepository
	ts   TicketService // навсякий
}

func NewMessageService(repo in_memory.MessageRepository, ts TicketService) MessageService {
	return &messageService{repo: repo, ts: ts}
}

func (s *messageService) NewMessage(ticketID string, mess message.Message) (bool, error) {
	err := message.ValidateMessageBody(mess.Body)

	if err != nil {
		return false, fmt.Errorf("empty body ")
	}
	_, err0 := s.ts.GetTicket(ticketID) // пока расплывчито что надо с самим тиктом делать если он существует
	if err0 != nil {
		s.ts.NewTicket(ticket.Ticket{
			ID:                    ticketID,
			RequesterID:           mess.SenderID,
			AgentID:               " ",
			Title:                 "placeholder", //тут наверное первые 25 символов как тайтл возьмем так как юзер просто заходит в чат поддержки и вводит сообщение или  сделаем по темам но это уже потом
			Description:           "placeholder",
			Status:                ticket.StatusWaiting,
			Version:               "placeholder",
			CreatedAt:             "placeholder",
			UpdatedAt:             "placeholder",
			ClosedAt:              "placeholder",
			TimeToLive:            "placeholder",
			ClosedReason:          "placeholder",
			CloseRequestedAt:      "placeholder",
			CloseRequestMessageID: "placeholder",
		})
		err1 := s.repo.StoreMessage(ticketID, mess)
		if err1 != nil {
			return false, fmt.Errorf("error with store message")
		}
		return true, nil
	}

	err1 := s.repo.StoreMessage(ticketID, mess)
	if err1 != nil || err != nil {
		return false, fmt.Errorf("error with store message")
	}

	return false, nil
}
func (s *messageService) LoadMessage(ticketID string) ([]message.Message, error) {

	return s.repo.LoadAllMessagesByTicketID(ticketID)
}
