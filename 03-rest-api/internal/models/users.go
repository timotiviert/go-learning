package models

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email" binding:"required,email"`
	Username     string `json:"username" binding:"required,alphanum,min=3,max=20"`
	PasswordHash string `json:"-"` // "-" means that it will never be returned.
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type RegisterUsers struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,alphanum,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}
