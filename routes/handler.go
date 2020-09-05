package routes

import (
	"sia/controllers"

	"github.com/gin-gonic/gin"
)

func Handler() {
	r := gin.Default()

	sia := r.Group("api/v1/")
	{
		sia.GET("users", controllers.GetUsers)
		sia.GET("akuns", controllers.GetAkuns)
		sia.GET("transaksis", controllers.GetTransaksi)
	}

	r.Run()
}
