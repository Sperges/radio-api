package models

type User struct {
	Id   int    `db:"Id" json:"id"`
	Name string `db:"Username" json:"username" binding:"required"`
}
