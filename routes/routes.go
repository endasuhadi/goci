package routes

import (
	"github.com/gin-gonic/gin"
	"app/controller"
)


func Routes(route *gin.Engine){

	private := route.Group("/api")
	private.GET("/", controller.Siswa)
	private.POST("/", controller.SaveSiswa)
	private.PUT("/:no", controller.UpdateSiswa)
	private.DELETE("/:no", controller.DeleteSiswa)

}