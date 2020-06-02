package repository

import "github.com/ssk0206/accountant/app/models"

type StudentRepositoryInterface interface {
	GetAll() []models.Student
	GetByRoomID(roomID int) (models.Student, bool)
	Create(student models.Student) (models.Student, error) // errorの実装に関しては疑問 status_codeを返すべき?
}
