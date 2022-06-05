package controllers

import (
	"database/sql"
	"net/http"
	"radio-api/models"
	"radio-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	err := services.CreateUser(user)

	switch err {
	case nil:
		c.JSON(http.StatusOK, "OK")
	case sql.ErrNoRows:
		_ = c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePrivate)
	default:
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}
}

func ReadUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	result, err := services.ReadUser(id)

	switch err {
	case nil:
		c.JSON(http.StatusOK, result)
	case sql.ErrNoRows:
		_ = c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePrivate)
	default:
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}
}

func ReadUsers(c *gin.Context) {
	if result, err := services.ReadUsers(); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	err := services.UpdateUser(user)

	switch err {
	case nil:
		c.JSON(http.StatusOK, "OK")
	case sql.ErrNoRows:
		_ = c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePrivate)
	default:
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return
	}

	err = services.DeleteUser(id)

	switch err {
	case nil:
		c.JSON(http.StatusOK, "OK")
	case sql.ErrNoRows:
		_ = c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePrivate)
	default:
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}
}
