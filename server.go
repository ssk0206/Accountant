package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/ssk0206/accountant/db"
	"github.com/ssk0206/accountant/router"

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

	r := router.NewRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
