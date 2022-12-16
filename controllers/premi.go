package controllers

import (
	"net/http"
	"pratice/database"
	"pratice/repository"
	"pratice/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPremi(c *gin.Context) {
	var (
		result gin.H
	)

	premis, err := repository.GetAllPremi(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": premis,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPremi(c *gin.Context) {
	var premi structs.Premi

	err := c.ShouldBindJSON(&premi)
	if err != nil {
		panic(err)
	}
	err = repository.InsertPremi(database.DbConnection, premi)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Premi",
	})

}

func UpdatePremi(c *gin.Context) {

	var premi structs.Premi

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&premi)
	if err != nil {
		panic(err)
	}
	premi.Id = int64(id)

	err = repository.UpdatePremi(database.DbConnection, premi)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Premi",
	})

}

func DeletePremi(c *gin.Context) {
	var premi structs.Premi

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	premi.Id = int64(id)

	err = repository.DeletePremi(database.DbConnection, premi)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Premi",
	})

}
