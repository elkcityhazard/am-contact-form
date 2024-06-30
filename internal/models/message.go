package models

import "time"

type Message struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Version        int       `json:"version"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	MessageContent string    `json:"message_content"`
	Token          string    `json:"token"`
	IP             string    `json:"ip_address"`
}
