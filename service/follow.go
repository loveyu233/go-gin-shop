package service

import (
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"
	"strconv"
	"time"
)

type Follow struct {
}

func (Follow) Follow(blogUserId string, isFollow bool, userId int) dto.Response {
	blogUserIdInt, _ := strconv.Atoi(blogUserId)
	follow := tb.TbFollow{
		UserID:       uint64(userId),
		FollowUserID: uint64(blogUserIdInt),
		CreateTime:   time.Now(),
	}
	// true 关注
	if isFollow {
		affected := global.MysqlDb.Model(&tb.TbFollow{}).Create(&follow).RowsAffected
		if affected > 0 {
			global.RedisDb.SAdd(global.Content, "follow:"+strconv.Itoa(userId), blogUserId)
			return dto.OkData("关注成功")
		}
		return dto.OkData("关注失败")
	}
	// false 取关
	affected := global.MysqlDb.Exec("delete from tb_follow where user_id = ? and follow_user_id = ?", follow.UserID, follow.FollowUserID).RowsAffected
	global.RedisDb.SRem(global.Content, "follow:"+strconv.Itoa(userId), blogUserId)
	if affected <= 0 {
		return dto.OkData("取关失败")
	}
	return dto.OkData("取消关注成功")
}

func (Follow) FollowOrNot(blogUserId string, userId int) dto.Response {
	result, _ := global.RedisDb.SIsMember(global.Content, "follow:"+strconv.Itoa(userId), blogUserId).Result()
	return dto.OkData(result)
}

func (Follow) Common(blogId string, userId int) dto.Response {
	result, _ := global.RedisDb.SInter(global.Content, "follow:"+blogId, "follow:"+strconv.Itoa(userId)).Result()
	var users []tb.TbUser
	global.MysqlDb.Model(&tb.TbUser{}).Where("id IN ?", result).Find(&users)
	for i := range users {
		users[i].Phone = ""
		users[i].Password = ""
	}
	return dto.OkData(users)
}
