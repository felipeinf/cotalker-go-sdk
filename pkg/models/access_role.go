package models

import "time"

type AccessRole struct {
	ID          string    `json:"_id"`
	Active      bool      `json:"active"`
	Company     string    `json:"company"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	ModifiedAt  time.Time `json:"modifiedAt"`
	Name        string    `json:"name"`
	Permissions []string  `json:"permissions"`
}
