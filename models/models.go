package models

type User struct {
	Id   int    `db:"Id" json:"id"`
	Name string `db:"Username" json:"username" binding:"required"`
}

type Radio struct {
	ID     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	UserId int    `json:"userId" binding:"required"`
}

type Playlist struct {
	ID     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	UserId int    `json:"userId" binding:"required"`
}
