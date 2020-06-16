package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssk0206/accountant/app/models"
	"github.com/ssk0206/accountant/app/repository"
)

func GetBills(c *gin.Context) {
	repo := repository.NewBillRepo()
	period := c.Query("period")
	bills, err := repo.GetAll(period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, bills)
}

func CreateBill(c *gin.Context) {
	repo := repository.NewBillRepo()

	type period struct {
		Period string `json:"period"`
	}
	var p period
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := repo.CreateBillPage(p.Period); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, nil)
}

func UpdateBill(c *gin.Context) {
	repo := repository.NewBillRepo()

	r := make([]models.Bill, 0)
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := repo.Update(r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, nil)
}
