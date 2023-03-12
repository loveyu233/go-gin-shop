package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/res"
	"go-gin-shop/enter/sys"
	"go-gin-shop/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type CommentApi struct {
}

var commentService = service.EnterServicesApp.CommentService

func (CommentApi) AddComment(c *gin.Context) {
	commentDto := &dto.CommentDTO{}
	err := c.Bind(commentDto)
	if err != nil {
		res.Err(c, err.Error())
		return
	}
	comment := dto.CommentDtoToComment(*commentDto, c.GetUint("userid"))
	id, err := commentService.InsertComment(comment)
	if err != nil {
		res.Err(c, "评论添加失败")
		return
	}
	res.OkData(c, id)
}
func (CommentApi) DelComment(c *gin.Context) {
	commentid, _ := primitive.ObjectIDFromHex(c.GetString("commentid"))
	err := commentService.DeleteComment(commentid)
	if err != nil {
		res.Err(c, "评论删除失败")
		return
	}
	res.Ok(c)
}

func (CommentApi) SelComment(c *gin.Context) {
	commentid, _ := primitive.ObjectIDFromHex(c.GetString("commentid"))
	comment, err := commentService.SelectCommentByID(commentid)
	if err != nil {
		res.Err(c, "查询失败")
		return
	}
	commentVO := dto.ToCommentVO([]sys.Comment{comment})
	res.OkData(c, commentVO)
}
func (CommentApi) SelCommentList(c *gin.Context) {
	bid, _ := strconv.Atoi(c.Query("bid"))
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	commentList, err := commentService.SelectCommentList(bid, page, pagesize)
	if err != nil {
		res.Err(c, "查询失败")
		return
	}
	commentVos := dto.ToCommentVO(commentList)
	res.OkData(c, commentVos)
}
func (CommentApi) AddReply(c *gin.Context) {
	replyDto := &dto.ReplyDTO{}
	c.Bind(replyDto)
	replyModel := dto.ReplyDtoToReply(*replyDto, c.GetUint("userid"))
	replyId, err := commentService.InsertReply(replyDto.ParentID, replyModel)
	if err != nil {
		res.Err(c, "添加回复失败")
		return
	}
	res.OkData(c, replyId)
}
func (CommentApi) DelReply(c *gin.Context) {
	commentid, _ := primitive.ObjectIDFromHex(c.Query("commentid"))
	err := commentService.DeleteComment(commentid)
	if err != nil {
		res.Err(c, "评论删除失败")
		return
	}
	res.Ok(c)
}
func (CommentApi) SelReply(c *gin.Context) {
	replyid, _ := primitive.ObjectIDFromHex(c.Query("replyid"))
	commentid, _ := primitive.ObjectIDFromHex(c.Query("commentid"))
	err := commentService.DeleteReply(commentid, replyid)
	if err != nil {
		res.Err(c, "回复删除失败")
		return
	}
	res.Ok(c)
}
