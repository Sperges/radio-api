package models

type Playlist struct {
	Id       int    `db:"Id" json:"id"`
	Name     string `db:"PlaylistName" json:"name" binding:"required"`
	UserId   int    `db:"UserId" json:"userId"`
	RadioIds []int  `json:"radioIds"`
}
