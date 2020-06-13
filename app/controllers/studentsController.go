package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssk0206/accountant/app/models"
	"github.com/ssk0206/accountant/app/repository"
)

func GetAllStudents(c *gin.Context) {
	repo := repository.NewStudentRepo()
	students := repo.GetAll()
	c.JSON(http.StatusOK, students)
}

func ShowStudent(c *gin.Context) {
	repo := repository.NewStudentRepo()
	roomid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, ok := repo.GetByRoomID(roomid)
	if ok == false {
		c.JSON(http.StatusNotFound, student)
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	repo := repository.NewStudentRepo()

	newStudent := models.Student{}
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student, err := repo.Create(newStudent)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	repo := repository.NewStudentRepo()

	student := models.Student{}
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repo.Delete(models.Student{}, student.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
