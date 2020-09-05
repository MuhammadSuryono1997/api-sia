package controllers

import (
	"net/http"
	"sia/config"
	"sia/models"

	"github.com/gin-gonic/gin"
)

func GetAkuns(c *gin.Context) {
	var akun []models.Akun
	var response models.ResponseSync

	config.GormDb.Find(&akun)
	response.Status = true
	response.Message = "Data berhasil di get"
	response.Data = akun
	c.JSON(http.StatusOK, response)
}
