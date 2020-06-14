package repository

import (
	"github.com/ssk0206/accountant/app/models"
	"github.com/ssk0206/accountant/db"
)

type studentRepo struct {
}

func NewStudentRepo() StudentRepositoryInterface {
	return &studentRepo{}
}

func (s *studentRepo) GetAll() []models.Student {
	students := make([]models.Student, 0)
	db.Db.Find(&students)
	return students
}

func (s *studentRepo) GetByRoomID(roomID int) (models.Student, bool) {
	student := models.Student{}
	if db.Db.Find(&student, "id=?", roomID).RecordNotFound() {
		return student, false
	}
	return student, true
}

func (s *studentRepo) Create(student models.Student) (models.Student, error) {
	if err := db.Db.Create(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (s *studentRepo) Delete(student models.Student, roomID string) error {
	if err := db.Db.Where("id = ?", roomID).Delete(&student).Error; err != nil {
		return err
	}
	return nil
}
