package tb

import "time"

type TbBlogComments struct {
	ID         uint64    `json:"id" gorm:"column:id"`                   // 主键
	UserID     uint64    `json:"user_id" gorm:"column:user_id"`         // 用户id
	BlogID     uint64    `json:"blog_id" gorm:"column:blog_id"`         // 探店id
	ParentID   uint64    `json:"parent_id" gorm:"column:parent_id"`     // 关联的1级评论id，如果是一级评论，则值为0
	AnswerID   uint64    `json:"answer_id" gorm:"column:answer_id"`     // 回复的评论id
	Content    string    `json:"content" gorm:"column:content"`         // 回复的内容
	Liked      uint      `json:"liked" gorm:"column:liked"`             // 点赞数
	Status     uint8     `json:"status" gorm:"column:status"`           // 状态，0：正常，1：被举报，2：禁止查看
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"` // 更新时间
}

func (m *TbBlogComments) TableName() string {
	return "tb_blog_comments"
}
