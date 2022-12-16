package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"pratice/controllers"
	"pratice/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func basicAuth(ctx *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := ctx.Request.BasicAuth()
	if hasAuth && user == "admin" && password == "password" {
		ctx.Next()
	} else if hasAuth && user == "editor" && password == "secret" {
		ctx.Next()
	} else {
		ctx.Abort()
		ctx.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"Error:": "Silahkan Masukan user dan password yang benar",
		})

		return
	}
}

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	//route
	router := gin.Default()

	//start route premi
	router.GET("/premi", basicAuth, controllers.GetAllPremi)
	router.POST("/premi", basicAuth, controllers.InsertPremi)
	router.PUT("/premi/:id", basicAuth, controllers.UpdatePremi)
	router.DELETE("/premi/:id", basicAuth, controllers.DeletePremi)

	//start route master
	router.GET("/master", basicAuth, controllers.GetAllMaster)
	router.POST("/master", basicAuth, controllers.InsertMaster)
	router.PUT("/master/:nik", basicAuth, controllers.UpdateMaster)
	router.DELETE("/master/:nik", basicAuth, controllers.DeleteMaster)

	//start route data kesehatan
	router.GET("/datakesehatan", basicAuth, controllers.GetAllDataKesehatan)
	router.GET("/datakesehatan/bynik/:nik", basicAuth, controllers.GetDataKesehatanByNIK)
	router.POST("/datakesehatan", basicAuth, controllers.InsertDataKesehatan)
	router.PUT("/datakesehatan/:nik", basicAuth, controllers.UpdateFaskes)

	//start route pembayaran
	router.GET("/datapembayaran/:nik/:periode", basicAuth, controllers.GetDataPembayaranByPeriode)
	router.POST("/bayar", basicAuth, controllers.Bayar)

	//start route tagihan
	router.GET("/tagihan/:nik", basicAuth, controllers.Tagihan)
	router.Run()
}
