package models

type GoogleProfile struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UID       string `json:"uid"`
	// we don't need fields below for test, but it's better to keep them in real life
	// Name string `json:"name"`
	// Email string `json:"email"`
}
