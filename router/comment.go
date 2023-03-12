package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
	"go-gin-shop/middleware"
)

type CommentRouter struct {
}

func (CommentRouter) InitCommentRouter(router *gin.RouterGroup) {
	commentApi := api.ApiGroupApp.CommentApi
	router.GET("/comment/get", commentApi.SelComment)
	router.GET("/comment/getlist", commentApi.SelCommentList)
	comment := router.Group("/commnt")
	comment.Use(middleware.Auth)
	{
		comment.POST("/add", commentApi.AddComment)
		comment.GET("/del", commentApi.DelComment)

	}
	reply := router.Group("/reply")
	reply.Use(middleware.Auth)
	{
		reply.POST("/add", commentApi.AddReply)
		reply.GET("/del", commentApi.DelReply)
		reply.GET("/sel", commentApi.SelReply)
	}
}
