package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ssk0206/accountant/app/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/students", controllers.GetAllStudents)
	router.GET("/students/:roomid", controllers.ShowStudent)
	router.POST("/students", controllers.CreateStudent)
	router.DELETE("/students", controllers.DeleteStudent)

	return router
}
