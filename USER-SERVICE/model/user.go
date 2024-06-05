package model

type User struct {
	ID string `json:"id"`
	Data string `json:"data"`
}

type Cation struct {
	UserID string `json:"user_id"`
	Action string `json:"action"`
}