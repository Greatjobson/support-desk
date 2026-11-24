package session

type Session struct {
	TokenHash  string `json:"token_hash"`
	UserID     string `json:"user_id"`
	TimeToLive string `json:"time_to_live"`
}
