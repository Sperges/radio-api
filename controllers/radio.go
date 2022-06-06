package controllers

import (
	"radio-api/models"
	"radio-api/services"

	"github.com/gin-gonic/gin"
)

func CreateRadio(c *gin.Context) {
	var radio models.Radio
	if err := BindStruct(c, &radio); err != nil {
		return
	}
	result, err := services.CreateRadio(radio)
	StructResponse(c, result, err)
}

func ReadRadio(c *gin.Context) {

}

func ReadRadios(c *gin.Context) {

}

func ReadRadioByIDFromUser(c *gin.Context) {

}

func ReadRadioByNameFromUser(c *gin.Context) {

}

func UpdateRadio(c *gin.Context) {

}

func DeleteRadio(c *gin.Context) {

}
