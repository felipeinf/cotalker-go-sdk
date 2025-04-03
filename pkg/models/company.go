package models

import "time"

type Company struct {
	ID          string    `json:"_id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Display     string    `json:"display"`
	Description string    `json:"description,omitempty"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	ModifiedAt  time.Time `json:"modifiedAt"`
}
