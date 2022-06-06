package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BindStruct[T any](c *gin.Context, obj T) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return err
	}
	return nil
}

func GetPathParamAsInt(c *gin.Context, identifier string) (int, error) {
	id, err := strconv.Atoi(c.Param(identifier))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypePrivate)
		return 0, err
	}
	return id, nil
}

func OkResponse(c *gin.Context, err error) {
	switch err {
	case nil:
		c.Status(http.StatusOK)
	case sql.ErrNoRows:
		_ = c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePrivate)
	default:
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}
}

func StructResponse(c *gin.Context, result any, err error) {
	switch err {
	case nil:
		c.JSON(http.StatusOK, result)
	case sql.ErrNoRows:
		_ = c.AbortWithError(http.StatusNotFound, err).SetType(gin.ErrorTypePrivate)
	default:
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}
}
