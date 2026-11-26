package in_memory

import (
	"fmt"

	"github.com/Greatjobson/support-desk/internal/model/message"
	"github.com/Greatjobson/support-desk/internal/model/ticket"
)

type TicketRepository interface {
	StoreTicket(ticketID string, data ticket.Ticket) error
	LoadTicket(ticketID string) (ticket.Ticket, error)
}

type TicketIMStorage struct {
	// знаю что тут надо какой то мютекс делать с дефером но мне лень темболее буду делать сразу в постгресс после того как налажу бизнес логику так что пусть остается
	ticketStore map[string]ticket.Ticket
}

func NewTicketIMStorage() *TicketIMStorage {
	return &TicketIMStorage{
		ticketStore: make(map[string]ticket.Ticket),
	}
}

func (s *TicketIMStorage) StoreTicket(ticketID string, data ticket.Ticket) error {
	s.ticketStore[ticketID] = data
	return nil
}

func (s *TicketIMStorage) LoadTicket(ticketID string) (ticket.Ticket, error) {

	value, ok := s.ticketStore[ticketID]
	if ok {
		return value, nil
	}
	return ticket.Ticket{}, fmt.Errorf("no such ticket: %s", ticketID)
}

type MessageRepository interface {
	StoreMessage(ticketID string, message message.Message) error
	LoadAllMessagesByTicketID(ticketID string) ([]message.Message, error)
}

type MessageIMStorage struct {
	messageStore map[string][]message.Message
}

func NewMessageIMStorage() *MessageIMStorage {
	return &MessageIMStorage{
		messageStore: make(map[string][]message.Message),
	}
}

// в бизнес написать как нью месседж
func (s *MessageIMStorage) StoreMessage(ticketID string, message message.Message) error {
	s.messageStore[ticketID] = append(s.messageStore[ticketID], message)
	return nil
}

func (s *MessageIMStorage) LoadAllMessagesByTicketID(ticketID string) ([]message.Message, error) {

	value, ok := s.messageStore[ticketID]
	if ok {
		return value, nil
	}
	return []message.Message{}, fmt.Errorf("no such message in ticket: %s", ticketID)
}
