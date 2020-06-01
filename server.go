package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/ssk0206/accountant/app/controllers"
	"github.com/ssk0206/accountant/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbuser string
	dbpass string
	dbname string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env")
	}

	dbuser = os.Getenv("DB_USER")
	dbpass = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB")
}

func main() {
	db.GormConnect(dbuser, dbpass, dbname)
	defer db.Close()

	log.Fatalln(serve())
}

func serve() error {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/students", controllers.GetAllStudents)
	router.GET("/students/:roomid", controllers.ShowStudent)

	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
