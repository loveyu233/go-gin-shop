package tb

import "time"

type TbFollow struct {
	ID           int64     `json:"id" gorm:"column:id"`                         // 主键
	UserID       uint64    `json:"user_id" gorm:"column:user_id"`               // 用户id
	FollowUserID uint64    `json:"follow_user_id" gorm:"column:follow_user_id"` // 关联的用户id
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`       // 创建时间
}

func (m *TbFollow) TableName() string {
	return "tb_follow"
}
