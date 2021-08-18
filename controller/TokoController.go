package controller

import (
	"api-apotek/data/model"
	"api-apotek/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokoController struct {
	tokoM model.TokoModel
}

func (e *TokoController) GetAllToko(c *gin.Context) {
	toko, err := e.tokoM.GetToko()
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.Success(c, toko)
}
