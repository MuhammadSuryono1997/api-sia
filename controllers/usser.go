package controllers

import (
	"net/http"
	"sia/config"
	"sia/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var user []models.User
	var response models.ResponseSync

	config.GormDb.Find(&user)
	response.Status = true
	response.Message = "Data berhasil di get"
	response.Data = user
	c.JSON(http.StatusOK, response)
}
