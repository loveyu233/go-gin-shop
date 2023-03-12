package tb

import "time"

type TbUser struct {
	ID         uint64    `json:"id,string" gorm:"column:id"`            // 主键
	Phone      string    `json:"phone" gorm:"column:phone"`             // 手机号码
	Password   string    `json:"password" gorm:"column:password"`       // 密码，加密存储
	NickName   string    `json:"nickName" gorm:"column:nick_name"`      // 昵称，默认是用户id
	Icon       string    `json:"icon" gorm:"column:icon"`               // 人物头像
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
}

func (m *TbUser) TableName() string {
	return "tb_user"
}
