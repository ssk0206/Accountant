package db

import (
	"log"

	"github.com/jinzhu/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func GormConnect(dbuser, dbpass, dbname string) {
	DBMS := "mysql"
	PROTOCOL := "tcp(mysql:3306)"
	CONNECT := dbuser + ":" + dbpass + "@" + PROTOCOL + "/" + dbname + "?charset=utf8&parseTime=true"
	Db, err = gorm.Open(DBMS, CONNECT)
	Db.LogMode(true)
	if err != nil {
		log.Fatalln(err)
	}
}

func Close() {
	if err = Db.Close(); err != nil {
		panic(err)
	}
}
