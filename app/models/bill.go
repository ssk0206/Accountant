package models

import "time"

type Bill struct {
	ID            uint       `json:"id"`
	RoomID        string     `json:"room_id"`
	Period        string     `json:"period"`
	PreMeterValue float64    `json:"pre_meter_val"`
	NewMeterValue float64    `json:"new_meter_val"`
	Bill          float64    `json:"bill"`
	AdditionalFee float64    `json:"additional_fee"`
	Remark        string     `json:"remark"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

type Response struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	PreMeterValue float64 `json:"pre-meter-value"`
	NewMeterValue float64 `json:"new-meter-value"`
	Bill          float64 `json:"bill"`
	AdditionalFee float64 `json:"additional_fee"`
	Remark        string  `json:"remark"`
	Period        string  `json:"period"`
}
