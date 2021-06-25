package utils

import (
	"api-apotek/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	res := model.ResponseGeneric{
		Status:  "Success",
		Message: "Loaded",
		Data:    data,
	}
	c.JSON(http.StatusOK, res)
}

func HandleError(c *gin.Context, status int, message string) {
	res := model.ResponseGeneric{
		Status:  "Failed",
		Message: message,
	}
	c.AbortWithStatusJSON(status, res)
}
