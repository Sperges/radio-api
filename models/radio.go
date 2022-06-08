package models

type Radio struct {
	Id          int    `db:"Id" json:"id"`
	Name        string `db:"RadioName" json:"name" binding:"required"`
	StreamUrl   string `db:"StreamUrl" json:"streamUrl" binding:"required"`
	UserId      int    `db:"UserId" json:"userId"`
	PlaylistIds []int  `json:"playlistIds"`
}
