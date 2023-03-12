package vo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentVo struct {
	ID        primitive.ObjectID `json:"id"`
	Content   string             `json:"content"`
	Author    UserVo             `json:"author"`
	Reply     []ReplyVo          `json:"reply"`
	CreatedAt int64              `json:"created_at"`
	At        []string           `json:"at"`
}

type ReplyVo struct {
	ID        primitive.ObjectID `json:"id"`
	Content   string             `json:"content"`
	Author    UserVo             `json:"author"`
	CreatedAt int64              `json:"created_at"`
	At        []string           `json:"at"`
}
