package repository

import "github.com/ssk0206/accountant/app/models"

type StudentRepositoryInterface interface {
	GetAll() []models.Student
	GetByRoomID(roomID int) (models.Student, bool)
}
