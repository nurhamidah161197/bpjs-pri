package controllers

import (
	"net/http"
	"pratice/database"
	"pratice/repository"
	"pratice/structs"

	"github.com/gin-gonic/gin"
)

func GetDataPembayaranByPeriodeNIK(c *gin.Context) {
	var datapembayaran structs.DataPembayaran

	nik := c.Param("nik")
	periode := c.Param("periode")

	datapembayaran.NIK = nik
	datapembayaran.Periode = periode

	datapembayarans, err := repository.GetDataPembayaranByPeriodeNIK(database.DbConnection, datapembayaran)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, datapembayarans)
}

func Bayar(c *gin.Context) {
	var datapembayaran structs.DataPembayaran

	err := c.ShouldBindJSON(&datapembayaran)
	if err != nil {
		panic(err)
	}

	err = repository.Bayar(database.DbConnection, datapembayaran)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Data Pembayaran",
	})

}
