package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ssk0206/Accountant/server/controllers"
)

func main() {
	serve()
}

func serve() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/students", controllers.GetAllStudents)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln("Failed to run serve: ", err)
	}
}
