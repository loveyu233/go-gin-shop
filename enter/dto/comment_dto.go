package dto

import (
	"go-gin-shop/enter/sys"
	"go-gin-shop/enter/vo"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentDTO struct {
	Bid     uint     `json:"bid,omitempty"`
	Content string   `json:"content,omitempty"`
	At      []string `json:"at,omitempty"`
}

type ReplyDTO struct {
	Bid          uint               `json:"bid,omitempty"`
	Content      string             `json:"content,omitempty"`
	ParentID     primitive.ObjectID `json:"parentID,omitempty"`
	At           []string           `json:"at,omitempty"`
	ReplyUserID  uint               `json:"replyUserID,omitempty"`
	ReplyContent string             `json:"replyContent,omitempty"`
}

type DeleteReplyDTO struct {
	CommentID primitive.ObjectID `json:"commentID,omitempty"`
	ReplyID   primitive.ObjectID `json:"replyID,omitempty"`
}

/**
 * 评论DTO结构体转化为Comment结构体
 * param: commentDTO 评论DTO结构体
 * return: comment结构体
 */
func CommentDtoToComment(commentDTO CommentDTO, userId uint) sys.Comment {
	return sys.Comment{
		ID:        primitive.NewObjectID(),
		Bid:       commentDTO.Bid,
		CreatedAt: time.Now().UnixMilli(),
		Content:   commentDTO.Content,
		Uid:       userId,
		Reply:     []sys.Reply{},
		At:        commentDTO.At,
		IsDelete:  false,
	}
}

/**
 * 回复DTO结构体转化为Reply结构体
 * param: replyDTO 回复DTO结构体
 * return: reply结构体
 */
func ReplyDtoToReply(replyDTO ReplyDTO, userId uint) sys.Reply {
	return sys.Reply{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now().UnixMilli(),
		Content:   replyDTO.Content,
		Uid:       userId,
		At:        replyDTO.At,
		IsDelete:  false,
	}
}
func ToCommentVO(comments []sys.Comment) []vo.CommentVo {
	length := len(comments)
	newComments := make([]vo.CommentVo, length)
	for i := 0; i < length; i++ {
		newComments[i].ID = comments[i].ID
		newComments[i].Content = comments[i].Content
		newComments[i].CreatedAt = comments[i].CreatedAt
		newComments[i].Reply = ToReplyVO(comments[i].Reply)
		newComments[i].At = comments[i].At
	}

	return newComments
}

func ToReplyVO(replies []sys.Reply) []vo.ReplyVo {
	length := len(replies)
	newReplies := make([]vo.ReplyVo, length)
	for i := 0; i < length; i++ {
		newReplies[i].ID = replies[i].ID
		newReplies[i].Content = replies[i].Content
		newReplies[i].CreatedAt = replies[i].CreatedAt
		newReplies[i].At = replies[i].At
	}
	return newReplies
}
