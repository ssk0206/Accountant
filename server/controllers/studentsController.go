package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssk0206/accountant/server/models"
)

func GetAllStudents(c *gin.Context) {

	student1 := models.Student{ID: 1, Name: "あいかわ"}
	student2 := models.Student{ID: 2, Name: "いがわ"}
	students := []models.Student{
		student1, student2,
	}

	c.JSON(200, students)
}
