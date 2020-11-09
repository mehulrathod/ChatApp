package api

import "time"

type UserResponse struct {
	ID        uint64    `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email,omitempty"`
	Mobile    string    `json:"mobile,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
