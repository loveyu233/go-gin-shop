package tb

import "time"

type TbShop struct {
	ID         uint64    `json:"id" gorm:"column:id"`                   // 主键
	Name       string    `json:"name" gorm:"column:name"`               // 商铺名称
	TypeID     uint64    `json:"type_id" gorm:"column:type_id"`         // 商铺类型的id
	Images     string    `json:"images" gorm:"column:images"`           // 商铺图片，多个图片以,隔开
	Area       string    `json:"area" gorm:"column:area"`               // 商圈，例如陆家嘴
	Address    string    `json:"address" gorm:"column:address"`         // 地址
	X          float64   `json:"x" gorm:"column:x"`                     // 经度
	Y          float64   `json:"y" gorm:"column:y"`                     // 维度
	AvgPrice   uint64    `json:"avg_price" gorm:"column:avg_price"`     // 均价，取整数
	Sold       uint      `json:"sold" gorm:"column:sold"`               // 销量
	Comments   uint      `json:"comments" gorm:"column:comments"`       // 评论数量
	Score      uint      `json:"score" gorm:"column:score"`             // 评分，1~5分，乘10保存，避免小数
	OpenHours  string    `json:"open_hours" gorm:"column:open_hours"`   // 营业时间，例如 10:00-22:00
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
	Dist       float64   `json:"distance" gorm:"-"`                     //商铺距离
}

func (m *TbShop) TableName() string {
	return "tb_shop"
}
