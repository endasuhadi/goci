package controller

import (
	"strconv"
	"app/model"
	"app/conf"
	"github.com/gin-gonic/gin"
)

var responseData model.ResponseData
var message string

func Siswa(c *gin.Context) {

	limit := c.DefaultQuery("limit", "0")

	i, _ := strconv.Atoi(limit)

	siswa := model.GetSiswa(i)

	if len(siswa) > 0 {
		message = conf.LoadSiswa
		responseData = model.ResponseData{conf.C200, message}
		Response := model.ResponseSiswa{ResponseData: responseData, Data: siswa}
		c.JSON(200, Response)
	}else{
		message = conf.NotFoundSiswa
		responseData = model.ResponseData{conf.C200, message}
		c.JSON(200, responseData)
	}
	
}

func SaveSiswa(c *gin.Context) {
	var siswa model.Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		message = "Error"
		responseData = model.ResponseData{conf.C403, message}
		c.JSON(403, responseData)
		return
	}
	message = "Sukses"
	model.SaveSiswa(&siswa)
	responseData = model.ResponseData{conf.C200, message}
	c.JSON(200, responseData)
}

func UpdateSiswa(c *gin.Context) {
	var siswa model.Siswa
	no, _ := strconv.Atoi(c.Param("no"))
	if err := c.ShouldBindJSON(&siswa); err != nil {
		message = "Error"
		responseData = model.ResponseData{conf.C403, message}
		c.JSON(403, responseData)
		return
	}

	message = "Sukses"
	model.UpdateSiswa(&siswa, no)
	responseData = model.ResponseData{conf.C200, message}
	c.JSON(200, responseData)
}

func DeleteSiswa(c *gin.Context) {
	no := c.Param("no")
	message = "Sukses"
	model.DeleteSiswa(no)
	responseData = model.ResponseData{conf.C200, message}
	c.JSON(200, responseData)
}