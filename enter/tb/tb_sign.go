package tb

import "time"

type TbSign struct {
	ID       uint64    `json:"id" gorm:"column:id"`               // 主键
	UserID   uint64    `json:"user_id" gorm:"column:user_id"`     // 用户id
	Year     int8      `json:"year" gorm:"column:year"`           // 签到的年
	Month    int8      `json:"month" gorm:"column:month"`         // 签到的月
	Date     time.Time `json:"date" gorm:"column:date"`           // 签到的日期
	IsBackup uint8     `json:"is_backup" gorm:"column:is_backup"` // 是否补签
}

func (m *TbSign) TableName() string {
	return "tb_sign"
}
