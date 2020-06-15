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

func (b *billRepo) GetAll(period string) ([]models.Response, error) {
	res := make([]models.Response, 0)
	if err := db.Db.Table("bills").Select("students.id, name, pre_meter_value, new_meter_value, bill, period, additional_fee, remark").Where("period = ?", period).Joins("left join students on students.id = bills.room_id").Scan(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (b *billRepo) CreateBillPage(period string) error {
	// INSERT INTO テーブル1(N1, N2, N3) SELECT C1, C2, C3 FROM テーブル2 WHERE 条件式;
	if err := db.Db.Exec("INSERT IGNORE INTO bills(room_id, period) SELECT id, " + "'" + period + "'" + " FROM students").Error; err != nil {
		return err
	}
	return nil
}
