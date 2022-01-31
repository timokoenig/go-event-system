package main

import "time"

type Event struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
	PublishedAt time.Time   `json:"published_at,omitempty"`
	StartedAt   time.Time   `json:"started_at,omitempty"`
	FinishedAt  time.Time   `json:"finished_at,omitempty"`
	Error       error       `json:"error,omitempty"`
}
