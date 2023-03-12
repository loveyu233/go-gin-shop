package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"
	"go-gin-shop/utils"
	"gorm.io/gorm"

	"strconv"
	"time"
)

// Blog 文章
type Blog struct {
}

// Hot 分页
func (Blog) Hot(page int, userId uint64) dto.Response {
	var blogs []tb.TbBlog
	global.MysqlDb.Model(&tb.TbBlog{}).Order("liked desc").Scopes(EnterServicesApp.PaginateService.paging(page)).Find(&blogs)
	if userId != 0 {
		for i := range blogs {
			user := EnterServicesApp.UserService.UserById(int(blogs[i].UserID))
			blogs[i].Name = user.NickName
			blogs[i].Icon = user.Icon
			result := utils.EnterUtilsApp.RedisBlogLiked.IsExist(global.RedisCacheLiked+strconv.Itoa(int(blogs[i].ID)), strconv.FormatUint(userId, 10))
			if result {
				blogs[i].IsLike = true
			}
		}
	}
	return dto.OkData(blogs)
}

// LikeBlog The detailed information:
// @Title likeBlog
// @Description 修改点赞数量
// @Param id
// @Return bool
func (Blog) LikeBlog(blogId, userId int) dto.Response {
	key := global.RedisCacheLiked + strconv.Itoa(blogId)
	result := utils.EnterUtilsApp.RedisBlogLiked.IsExist(key, strconv.Itoa(userId))
	// 不存在，点赞+1
	if !result {
		// 把数据库中对用的blog数据liked + 1
		affected := global.MysqlDb.Model(&tb.TbBlog{}).Scopes(EnterServicesApp.PaginateService.byId(blogId)).Update("liked", gorm.Expr("liked+1")).RowsAffected
		// 只有数据库更新成功才能修改redis数据
		if affected > 0 {
			// 把用户id添加到blog对应的redis set集合中
			utils.EnterUtilsApp.RedisBlogLiked.LikedAdd(key, userId)
		}
		return dto.Ok()
	}
	// 存在，取消点赞，点赞-1
	affected := global.MysqlDb.Model(&tb.TbBlog{}).Scopes(EnterServicesApp.PaginateService.byId(blogId)).Update("liked", gorm.Expr("liked-1")).RowsAffected
	// 只有数据库更新成功才能修改redis数据
	if affected > 0 {
		// 把blog对应的redis set集合中的userid删除掉
		utils.EnterUtilsApp.RedisBlogLiked.LikedDel(key, userId)
	}
	return dto.Ok()
}

// ById The detailed information:
// @Title ById
// @Description 查询具体blog信息
// @Param id
// @Return tb.TbBlog
func (Blog) ById(id int) dto.Response {
	blog := tb.TbBlog{}
	if global.MysqlDb.Model(&tb.TbBlog{}).Scopes(EnterServicesApp.PaginateService.byId(id)).Find(&blog).RowsAffected <= 0 {
		return dto.Err("blog不存在")
	}
	user := tb.TbUser{}
	global.MysqlDb.Model(&tb.TbUser{}).Where("id = ?", blog.UserID).First(&user)
	blog.Name = user.NickName
	blog.Icon = user.Icon
	result := utils.EnterUtilsApp.RedisBlogLiked.IsExist(global.RedisCacheLiked+strconv.Itoa(int(blog.ID)), strconv.FormatUint(user.ID, 10))
	if result {
		blog.IsLike = true
	}
	return dto.OkData(blog)
}

// OfMe The detailed information:
// @Title OfMe
// @Description 查询用户的blog
// @Param c
// @Return []tb.TbBlog
func (Blog) OfMe(c *gin.Context) dto.Response {
	user := utils.EnterUtilsApp.JwtUtils.JwtGetUser(c)
	var blogs []tb.TbBlog
	global.MysqlDb.Model(&tb.TbBlog{}).Where("user_id = ?", user.ID).Find(&blogs)
	for i := range blogs {
		if utils.EnterUtilsApp.RedisBlogLiked.IsExist(global.RedisCacheLiked+strconv.Itoa(int(blogs[i].ID)), strconv.FormatUint(utils.EnterUtilsApp.JwtUtils.JwtGetUser(c).ID, 10)) {
			blogs[i].IsLike = true
		}
	}
	return dto.OkData(blogs)
}

// Blog The detailed information:
// @Title Blog
// @Description 上传文章
// @Param blog
// @Param user
// @Return dto.Response
func (Blog) Blog(blog *tb.TbBlog, user *tb.TbUser) dto.Response {
	blog.UserID = user.ID
	// 保存blog
	if global.MysqlDb.Model(&tb.TbBlog{}).Create(blog).RowsAffected <= 0 {
		return dto.Err("blog保存失败")
	}
	var fans []tb.TbFollow
	// 查询粉丝
	global.MysqlDb.Model(&tb.TbFollow{}).Where("follow_user_id = ?", blog.UserID).Find(&fans)
	// 推送消息
	for i := range fans {
		key := global.RedisFeedFans + strconv.Itoa(int(fans[i].UserID))
		z := redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: blog.ID,
		}
		global.RedisDb.ZAdd(global.Content, key, &z)
	}
	return dto.OkData(blog.ID)
}

// BlogLikes The detailed information:
// @Title BlogLikes
// @Description 点赞排行榜
// @Param blogId
// @Return dto.Response
func (Blog) BlogLikes(blogId string) dto.Response {
	users := utils.EnterUtilsApp.RedisBlogLiked.GetBlogLikedUser(global.RedisCacheLiked + blogId)
	return dto.OkData(users)
}

func (Blog) OfUser(userId, current string) dto.Response {
	var blogs []tb.TbBlog
	currentInt, _ := strconv.Atoi(current)
	global.MysqlDb.Model(&tb.TbBlog{}).Where("user_id = ?", userId).Scopes(EnterServicesApp.PaginateService.paging(currentInt)).Find(&blogs)
	return dto.OkData(blogs)
}

func (Blog) OfFollowService(lastId, offset string, userId uint64) dto.Response {
	key := global.RedisFeedFans + strconv.Itoa(int(userId))
	offsetInt, _ := strconv.Atoi(offset)
	result, _ := global.RedisDb.ZRangeByScore(global.Content, key, &redis.ZRangeBy{
		Min:    "0",
		Max:    lastId,
		Offset: int64(offsetInt),
		Count:  5,
	}).Result()
	if len(result) == 0 {
		return dto.Ok()
	}
	var blogs []tb.TbBlog
	global.MysqlDb.Model(&tb.TbBlog{}).Where("id IN ?", result).Find(&blogs)
	for i := range blogs {
		blogUser := tb.TbUser{}
		global.MysqlDb.Model(&tb.TbUser{}).Where("id = ?", blogs[i].UserID).Find(&blogUser)
		blogs[i].UserID = blogUser.ID
		blogs[i].Icon = blogUser.Icon
	}
	// 索引为0的是最小的
	score, _ := global.RedisDb.ZScore(global.Content, key, result[0]).Result()
	// 查看sortSet中有几个score为lastId的值，而这个key对应值的个数即为byScore的offSet值
	count, _ := global.RedisDb.ZCount(global.Content, key, strconv.Itoa(int(score)), strconv.Itoa(int(score))).Result()
	return dto.OkData(dto.BlogFollow{
		List:    blogs,
		MinTime: score,
		Offset:  count,
	})
}
