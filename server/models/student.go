package models

type Student struct {
	ID            int     `json:"id"`
	RoomID        int     `json:"roomid"`
	Name          string  `json:"name"`
	PreMeterValue float64 `json:"pre_meter_val"`
	NewMeterValue float64 `json:"new_meter_val"`
	Bill          float64 `json:"bill"`
}
