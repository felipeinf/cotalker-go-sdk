package models

import "time"

type User struct {
	ID        string    `json:"_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Company   string    `json:"company"`
	IsActive  bool      `json:"isActive"`
	Roles     []string  `json:"roles,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
