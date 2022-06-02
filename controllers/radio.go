package controllers

import (
	"net/http"
	"radio-api/models"
	"radio-api/services"

	"github.com/gin-gonic/gin"
)

func CreateRadio(c *gin.Context) {
	var radio models.Radio
	if err := c.ShouldBindJSON(&radio); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.CreateRadio(radio); err != nil {
		_ = c.AbortWithError(http.StatusNotImplemented, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func ReadRadio(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.ReadRadio(id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func ReadRadioByIDFromUser(c *gin.Context) {

}

func ReadRadioByNameFromUser(c *gin.Context) {

}

func UpdateRadio(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	var radio models.Radio
	if err := c.ShouldBindJSON(&radio); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.UpdateRadio(id, radio); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func DeleteRadio(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if err := services.DeleteRadio(id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, "")
	}
}
