package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssk0206/accountant/app/repository"
)

func GetAllStudents(c *gin.Context) {
	repo := repository.NewStudentRepo()
	students := repo.GetAll()
	c.JSON(200, students)
}

func ShowStudent(c *gin.Context) {
	repo := repository.NewStudentRepo()
	roomid, err := strconv.Atoi(c.Param("roomid"))
	if err != nil {
		log.Fatalln("Failed to Atoi(): ", err)
	}

	student, ok := repo.GetByRoomID(roomid)
	if ok == false {
		c.JSON(404, student)
		return
	}
	c.JSON(200, student)
}

// func CreateStudent(c *gin.Context) {
// 	student1 := models.Student{Name: "赤城 武彦", RoomID: "101", PreMeterValue: 2000, NewMeterValue: 2010, Bill: 200}
// 	db.Db.Create(&student1)
// 	c.JSON(201, student1)
// }
