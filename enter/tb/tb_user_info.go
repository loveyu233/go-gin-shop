package tb

import "time"

type TbUserInfo struct {
	UserID     uint64    `json:"user_id" gorm:"column:user_id"`         // 主键，用户id
	City       string    `json:"city" gorm:"column:city"`               // 城市名称
	Introduce  string    `json:"introduce" gorm:"column:introduce"`     // 个人介绍，不要超过128个字符
	Fans       uint      `json:"fans" gorm:"column:fans"`               // 粉丝数量
	Followee   uint      `json:"followee" gorm:"column:followee"`       // 关注的人的数量
	Gender     uint8     `json:"gender" gorm:"column:gender"`           // 性别，0：男，1：女
	Birthday   time.Time `json:"birthday" gorm:"column:birthday"`       // 生日
	Credits    uint      `json:"credits" gorm:"column:credits"`         // 积分
	Level      uint8     `json:"level" gorm:"column:level"`             // 会员级别，0~9级,0代表未开通会员
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
}

func (m *TbUserInfo) TableName() string {
	return "tb_user_info"
}
