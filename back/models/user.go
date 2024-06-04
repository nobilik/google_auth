package models

type User struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"-"`          // for test is not encrypted
	IsGoogler bool   `json:"is_googler"` // non db field for disabling email field
}
