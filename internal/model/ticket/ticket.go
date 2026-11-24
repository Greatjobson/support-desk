package ticket

type Ticket struct {
	ID          string `json:"id"`
	RequesterID string `json:"requester_id"`
	AgentID     string `json:"agent_id"`

	Title       string `json:"title"` // из первых ну незнаю 25 слов получаем тайт по которому можно будет кликнуть
	Description string `json:"description"`
	Status      Status `json:"status"`
	Version     string `json:"version"`

	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	ClosedAt   string `json:"closed_at"`
	TimeToLive string `json:"time_to_live"`

	ClosedReason          string `json:"closed_reason"`
	CloseRequestedAt      string `json:"close_requested_at"`
	CloseRequestMessageID string `json:"close_request_message_id"`
}

type Status string

const (
	StatusWaiting    Status = "waiting"
	StatusInProgress Status = "in_progress"
	StatusResolved   Status = "resolved"
	StatusClosed     Status = "closed"
)
