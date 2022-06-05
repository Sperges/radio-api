package models

type Playlist struct {
	ID     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	UserId int    `json:"userId" binding:"required"`
}
