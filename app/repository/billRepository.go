package repository

import (
	"github.com/ssk0206/accountant/app/models"
	"github.com/ssk0206/accountant/db"
)

type billRepo struct {
}

func NewBillRepo() BillRepositoryInterface {
	return &billRepo{}
}

func (b *billRepo) GetAll(period string) ([]models.Bill, error) {
	bills := make([]models.Bill, 0)
	if err := db.Db.Where("period = ?", period).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (b *billRepo) CreateBillPage(period string) error {
	// INSERT INTO テーブル1(N1, N2, N3) SELECT C1, C2, C3 FROM テーブル2 WHERE 条件式;
	if err := db.Db.Exec("INSERT IGNORE INTO bills(room_id, period) SELECT id, " + "'" + period + "'" + " FROM students").Error; err != nil {
		return err
	}
	return nil
}
