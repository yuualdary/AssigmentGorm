package main

import (
	"Assigment2/config"
	"Assigment2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/mux"
)

func main() {

	router := gin.Default()

	config.ConnectDatabase()

	//route

	router.PUT("/update/:id", controllers.UpdateOrder)
	router.POST("/create", controllers.CreateOrder)
	router.GET("/order/:id", controllers.GetCurrentOrderById)

	router.Run(":3000")

}
