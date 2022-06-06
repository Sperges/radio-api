package controllers

import (
	"radio-api/models"
	"radio-api/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := BindStruct(c, &user); err != nil {
		return
	}
	result, err := services.CreateUser(user)
	StructResponse(c, result, err)
}

func ReadUser(c *gin.Context) {
	id, paramErr := GetPathParamAsInt(c, "id")
	if paramErr != nil {
		return
	}
	result, resultErr := services.ReadUser(id)
	StructResponse(c, result, resultErr)
}

func ReadUsers(c *gin.Context) {
	result, err := services.ReadUsers()
	StructResponse(c, result, err)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := BindStruct(c, &user); err != nil {
		return
	}
	OkResponse(c, services.UpdateUser(user))
}

func DeleteUser(c *gin.Context) {
	id, err := GetPathParamAsInt(c, "id")
	if err != nil {
		return
	}
	OkResponse(c, services.DeleteUser(id))
}
