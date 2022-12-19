package controllers

import (
	"net/http"
	"pratice/database"
	"pratice/repository"
	"pratice/structs"

	"github.com/gin-gonic/gin"
)

func GetAllDataKesehatan(c *gin.Context) {
	var (
		result gin.H
	)

	datakesehatans, err := repository.GetAllDataKesehatan(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": datakesehatans,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetDataKesehatanByNIK(c *gin.Context) {
	var datakesehatan structs.DataKesehatan

	nik := c.Param("nik")

	datakesehatan.NIK = nik

	results, err := repository.GetDataKesehatanByNIK(database.DbConnection, datakesehatan)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

func InsertDataKesehatan(c *gin.Context) {
	var datakesehatan structs.DataKesehatan

	err := c.ShouldBindJSON(&datakesehatan)
	if err != nil {
		panic(err)
	}

	err = repository.InsertDataKesehatan(database.DbConnection, datakesehatan)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Data Kesehatan",
	})

}

func UpdateDataKesehatan(c *gin.Context) {

	var datakesehatan structs.DataKesehatan

	nik := c.Param("nik")

	err := c.ShouldBindJSON(&datakesehatan)
	if err != nil {
		panic(err)
	}
	datakesehatan.NIK = nik

	err = repository.UpdateDataKesehatan(database.DbConnection, datakesehatan)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Data Kesehatan",
	})

}

func Tagihan(c *gin.Context) {
	var (
		result gin.H
	)
	var datakesehatan structs.DataKesehatan

	nik := c.Param("nik")

	datakesehatan.NIK = nik

	datakesehatans, total_tagihan := repository.Tagihan(database.DbConnection, datakesehatan)

	result = gin.H{
		"tagihan": total_tagihan,
		"result":  datakesehatans,
	}

	c.JSON(http.StatusOK, result)
}

func DeleteDataKesehatan(c *gin.Context) {
	var datakesehatan structs.DataKesehatan

	nik := c.Param("nik")

	datakesehatan.NIK = nik

	err := repository.DeleteDataKesehatan(database.DbConnection, datakesehatan)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Data Kesehatan",
	})

}
