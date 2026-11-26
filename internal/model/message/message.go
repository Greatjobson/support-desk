package message

// Message с ID, TicketID, SenderID (*int64), Kind, Body, CreatedA
type Message struct {
	ID            string `json:"id"`
	TiketID       string `json:"ticket_id"`
	SenderID      string `json:"sender_id"`
	MessageAuthor Author `json:"message_author"`
	Body          string `json:"body"`
	CreatedAt     string `json:"created_at"`
}

type Author string

const (
	User    Author = "user"
	Admin   Author = "admin" // на будущее
	Agent   Author = "agent"
	System  Author = "system"
	AIAgent Author = "AIAgent" // на будущее
)
