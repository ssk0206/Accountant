package repository

import "github.com/ssk0206/accountant/app/models"

type BillRepositoryInterface interface {
	GetAll(period string) ([]models.Bill, error)
	CreateBillPage(yearMonth string) error
}
