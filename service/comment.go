package service

import (
	"context"
	"errors"
	"go-gin-shop/enter/sys"
	"go-gin-shop/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
}

// 插入评论
func (Comment) InsertComment(comment sys.Comment) (primitive.ObjectID, error) {
	_, err := global.MongoDb.Comment().InsertOne(context.TODO(), comment)
	return comment.ID, err
}

// 插入回复
func (Comment) InsertReply(commentId primitive.ObjectID, reply sys.Reply) (primitive.ObjectID, error) {
	_, err := global.MongoDb.Comment().UpdateOne(context.TODO(), bson.M{"_id": commentId}, bson.M{
		"$addToSet": bson.M{
			"reply": reply,
		},
	})

	return reply.ID, err
}

// 查询评论
func (Comment) SelectCommentByID(commentId primitive.ObjectID) (sys.Comment, error) {
	var comments []sys.Comment
	cursor, err := global.MongoDb.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id": commentId,
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
			},
		},
		bson.M{
			"$limit": 1,
		},
	})

	if err != nil {
		return sys.Comment{}, err
	}

	if err := cursor.All(context.TODO(), &comments); err != nil {
		return sys.Comment{}, err
	}

	if len(comments) == 0 {
		return sys.Comment{}, errors.New("没有数据")
	}

	return comments[0], nil
}

// 查询回复
func (Comment) SelectReplyByID(commentId, replyId primitive.ObjectID) (sys.Reply, error) {
	var replies []sys.Reply
	cursor, err := global.MongoDb.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id": commentId,
			},
		},
		bson.M{
			"$project": bson.M{
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item._id", replyId},
						},
					},
				},
			},
		},
		bson.M{
			"$unwind": "$reply",
		},
		bson.M{
			"$project": bson.M{
				"_id":        "$reply._id",
				"uid":        "$reply.uid",
				"content":    "$reply.content",
				"created_at": "$reply.created_at",
			},
		},
		bson.M{
			"$limit": 1,
		},
	})

	if err != nil {
		return sys.Reply{}, err
	}

	if err := cursor.All(context.TODO(), &replies); err != nil {
		return sys.Reply{}, err
	}

	if len(replies) == 0 {
		return sys.Reply{}, errors.New("获取回复失败")
	}

	return replies[0], nil
}

// 查询评论
func (Comment) SelectCommentList(bid, page, pageSize int) ([]sys.Comment, error) {
	var comments []sys.Comment
	cursor, err := global.MongoDb.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"bid":       bid,
				"is_delete": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item.is_delete", false},
						},
					},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"uid":        "$uid",
				"content":    "$content",
				"created_at": "$created_at",
				"reply": bson.M{
					"$slice": bson.A{"$reply", 0, 2},
				},
			},
		},
		bson.M{
			"$skip": (page - 1) * pageSize,
		},
		bson.M{
			"$limit": pageSize,
		},
	})

	if err != nil {
		return comments, err
	}

	if err := cursor.All(context.TODO(), &comments); err != nil {
		return comments, err
	}

	return comments, nil
}

// 查询回复
func (Comment) SelectReplyList(id string, page, pageSize int) ([]sys.Reply, error) {
	var replies []sys.Reply
	objectId, _ := primitive.ObjectIDFromHex(id)
	cursor, err := global.MongoDb.Comment().Aggregate(context.TODO(), bson.A{
		bson.M{
			"$match": bson.M{
				"_id":       objectId,
				"is_delete": false,
			},
		},
		bson.M{
			"$project": bson.M{
				"reply": bson.M{
					"$filter": bson.M{
						"input": "$reply",
						"as":    "item",
						"cond": bson.M{
							"$eq": bson.A{"$$item.is_delete", false},
						},
					},
				},
			},
		},
		bson.M{
			"$unwind": "$reply",
		},
		bson.M{
			"$project": bson.M{
				"_id":        "$reply._id",
				"uid":        "$reply.uid",
				"content":    "$reply.content",
				"created_at": "$reply.created_at",
			},
		},
		bson.M{
			"$skip": (page-1)*pageSize + 2,
		},
		bson.M{
			"$limit": pageSize,
		},
	})

	if err != nil {
		return replies, err
	}

	if err := cursor.All(context.TODO(), &replies); err != nil {
		return replies, err
	}

	return replies, nil
}

func (Comment) DeleteComment(objectId primitive.ObjectID) error {
	_, err := global.MongoDb.Comment().UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.M{
		"$set": bson.M{
			"is_delete": true,
		},
	})

	return err
}

func (Comment) DeleteReply(commentId, replyId primitive.ObjectID) error {
	filter := bson.M{
		"_id":       commentId,
		"reply._id": replyId,
	}

	_, err := global.MongoDb.Comment().UpdateOne(context.TODO(), filter, bson.M{
		"$set": bson.M{
			"reply.$.is_delete": true,
		},
	})

	return err
}
