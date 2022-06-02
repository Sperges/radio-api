package models

type User struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
}

type Radio struct {
	ID     int    `json:"Id"`
	Name   string `json:"Name"`
	UserId int    `json:"UserId"`
}

type Playlist struct {
	ID     int    `json:"Id"`
	Name   string `json:"Name"`
	UserId int    `json:"UserId"`
}
