package controllers

import (
	"radio-api/models"
	"radio-api/services"

	"github.com/gin-gonic/gin"
)

func CreateRadio(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "user")
	if paramErr != nil {
		return
	}
	var radio models.Radio
	if bindErr := BindStruct(c, &radio); bindErr != nil {
		return
	}
	result, resultErr := services.CreateRadio(userId, radio)
	StructResponse(c, result, resultErr)
}

func ReadRadios(c *gin.Context) {
	userId, paramErr := GetPathParamAsInt(c, "user")
	if paramErr != nil {
		return
	}
	result, resultErr := services.ReadRadiosByUserId(userId)
	StructResponse(c, result, resultErr)
}

func UpdateRadio(c *gin.Context) {
	_, paramErr := GetPathParamAsInt(c, "user")
	if paramErr != nil {
		return
	}
	var radio models.Radio
	if err := BindStruct(c, &radio); err != nil {
		return
	}
	OkResponse(c, services.UpdateRadio(radio))
}

func DeleteRadio(c *gin.Context) {
	_, paramErr := GetPathParamAsInt(c, "user")
	if paramErr != nil {
		return
	}
	id, paramErr := GetPathParamAsInt(c, "id")
	if paramErr != nil {
		return
	}
	OkResponse(c, services.DeleteRadio(id))
}
