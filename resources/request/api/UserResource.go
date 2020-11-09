package api

import "time"

type SignUpRequest struct {
	ID        uint64
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Mobile    string    `json:"mobile" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
