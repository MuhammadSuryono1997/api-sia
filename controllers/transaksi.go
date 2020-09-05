package controllers

import (
	"net/http"
	"sia/config"
	"sia/models"

	"github.com/gin-gonic/gin"
)

var transaksi []models.Transaksi
var response models.ResponseSync

func GetTransaksi(c *gin.Context) {
	config.GormDb.Find(&transaksi)
	response.Status = true
	response.Message = "Data berhasil di get"
	response.Data = transaksi
	c.JSON(http.StatusOK, response)
}
