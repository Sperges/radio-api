package models

import "github.com/google/uuid"

type User struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
}

type Radio struct {
	ID     uuid.UUID `json:"Id"`
	Name   string    `json:"Name"`
	UserId int       `json:"UserId"`
}

type Playlist struct {
	ID     uuid.UUID `json:"Id"`
	Name   string    `json:"Name"`
	UserId int       `json:"UserId"`
}
