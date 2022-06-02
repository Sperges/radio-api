package controllers

import (
	"net/http"
	"radio-api/models"
	"radio-api/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.CreateUser(user); err != nil {
		_ = c.AbortWithError(http.StatusNotImplemented, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func ReadUser(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.ReadUser(id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func UpdateUser(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if result, err := services.UpdateUser(id, user); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func DeleteUser(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if err := services.DeleteUser(id); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, "")
	}
}
