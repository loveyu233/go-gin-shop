package tb

type TbBlog struct {
	ID         uint64 `json:"id" gorm:"column:id"`                   // 主键
	ShopID     int64  `json:"shopId" gorm:"column:shop_id"`          // 商户id
	UserID     uint64 `json:"userId" gorm:"column:user_id"`          // 用户id
	Title      string `json:"title" gorm:"column:title"`             // 标题
	Images     string `json:"images" gorm:"column:images"`           // 探店的照片，最多9张，多张以,隔开
	Content    string `json:"content" gorm:"column:content"`         // 探店的文字描述
	Liked      uint   `json:"liked" gorm:"column:liked"`             // 点赞数量
	Comments   uint   `json:"comments" gorm:"column:comments"`       // 评论数量
	CreateTime string `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime string `json:"update_time" gorm:"column:update_time"` // 更新时间
	Icon       string `json:"icon" gorm:"-"`                         //用户图标
	Name       string `json:"name" gorm:"-"`                         // 用户名称
	IsLike     bool   `json:"isLike" gorm:"-"`                       //是否点赞
}

func (m *TbBlog) TableName() string {
	return "tb_blog"
}
