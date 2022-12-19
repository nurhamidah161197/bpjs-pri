package controllers

import (
	"net/http"
	"pratice/database"
	"pratice/repository"
	"pratice/structs"

	"github.com/gin-gonic/gin"
)

func GetAllMaster(c *gin.Context) {
	var (
		result gin.H
	)

	masters, err := repository.GetAllMaster(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": masters,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertMaster(c *gin.Context) {
	var master structs.Master

	err := c.ShouldBindJSON(&master)
	if err != nil {
		panic(err)
	}

	err = repository.InsertMaster(database.DbConnection, master)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Master",
	})

}

func UpdateMaster(c *gin.Context) {

	var master structs.Master

	nik := c.Param("nik")

	err := c.ShouldBindJSON(&master)
	if err != nil {
		panic(err)
	}
	master.NIK = nik

	err = repository.UpdateMaster(database.DbConnection, master)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Master",
	})

}

func DeleteMaster(c *gin.Context) {
	var master structs.Master

	nik := c.Param("nik")

	master.NIK = nik

	err := repository.DeleteMaster(database.DbConnection, master)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Data Master & Kesehatan",
	})

}
