package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssk0206/accountant/app/models"
)

func GetAllStudents(c *gin.Context) {

	student1 := models.Student{ID: 1, Name: "宮城 リョータ", RoomID: 101, PreMeterValue: 2000, NewMeterValue: 2010, Bill: 200}
	student2 := models.Student{ID: 1, Name: "桜木 花道", RoomID: 102, PreMeterValue: 2000, NewMeterValue: 2010, Bill: 200}
	students := []models.Student{
		student1, student2,
	}

	c.JSON(200, students)
}
