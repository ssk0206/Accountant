package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ssk0206/Accountant/server/controllers"
)

var (
	dbuser string
	dbpass string
	dbname string
)

func init() {
	err := godotenv.Load(fmt.Sprintf("../.env"))
	if err != nil {
		log.Fatalln("Failed to load .env")
	}

	dbuser = os.Getenv("DB_USER")
	dbpass = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB")
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	// PROTOCOL := ""
	// CONNECT := dbuser + ":" + dbpass + "@" + PROTOCOL + "/" + dbname
	CONNECT := "root:@/accountant"
	db, err := gorm.Open(DBMS, CONNECT)
	db.LogMode(true)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()
	// mysql()
	log.Fatalln(serve())
}

func mysql() {
	db, err := sql.Open("mysql", "root:@/accountant")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される
}

func serve() error {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/students", controllers.GetAllStudents)

	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
