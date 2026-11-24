package user

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         Role   `json:"roles"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    string `json:"created_at"`
}

type Role string

const (
	Requester Role = "Requester"
	Agent     Role = "Agent"
	Admin     Role = "Admin" // ff
)
