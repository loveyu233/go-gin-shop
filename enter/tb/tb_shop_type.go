package tb

import "time"

type TbShopType struct {
	ID         uint64    `json:"id" gorm:"column:id"`                   // 主键
	Name       string    `json:"name" gorm:"column:name"`               // 类型名称
	Icon       string    `json:"icon" gorm:"column:icon"`               // 图标
	Sort       uint      `json:"sort" gorm:"column:sort"`               // 顺序
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
}

func (m *TbShopType) TableName() string {
	return "tb_shop_type"
}
