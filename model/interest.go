package model

import (
	"time"
)

type Interest struct {
	ID        string    `json:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user"`
	Apartment Apartment `json:"apartment"`
}
