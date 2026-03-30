package models

import "time"

type Job struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Priority  string    `json:"priority"`
	Attempts  int       `json:"attempts"`
	MaxRetry  int       `json:"max_retry"`
	CreatedAt time.Time `json:"created_at"`
}
