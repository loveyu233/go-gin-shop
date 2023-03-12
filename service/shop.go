package service

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"
	"go-gin-shop/redisLock"
	"go-gin-shop/utils"

	"strconv"
	"time"
)

// Shop 店铺
type Shop struct {
}

// ShowType The detailed information:
// @Title ShowType
// @Description 查询店铺分类信息
// @Return []tb.TbShopType
func (Shop) ShowType() dto.Response {
	result, err := global.RedisDb.Get(global.Content, global.ShopTypeListCache).Result()
	var st []tb.TbShopType
	if err != nil {
		global.MysqlDb.Model(&tb.TbShopType{}).Find(&st)
		marshal, _ := json.Marshal(st)
		global.RedisDb.Set(global.Content, global.ShopTypeListCache, string(marshal), global.TTLTime)
		return dto.OkData(st)
	}
	err = json.Unmarshal([]byte(result), &st)
	if err != nil {
		return dto.Err("json err")
	}
	return dto.OkData(st)
}

// OfType The detailed information:
// @Title OfType
// @Description shop根据typeId进行分页查询
// @Param id
// @Param current 页码
// @Return []tb.TbShop
func (Shop) OfType(id, current, x, y string) dto.Response {
	xFloat, _ := strconv.ParseFloat(x, 10)
	yFloat, _ := strconv.ParseFloat(y, 10)
	currentInt, _ := strconv.Atoi(current)
	searchRes := utils.EnterUtilsApp.RedisGeoUtil.SearchGeo(global.RedisShopGeo+id, xFloat, yFloat, currentInt)
	var idIn []string
	idDist := make(map[string]float64)
	for _, v := range searchRes {
		idIn = append(idIn, v.Name)
		idDist[v.Name] = v.Dist
	}
	var shops []tb.TbShop
	global.MysqlDb.Model(&tb.TbShop{}).Where("type_id=? and Id IN ?", id, idIn).Scopes(EnterServicesApp.PaginateService.paging(currentInt)).Find(&shops)
	for i := range shops {
		shops[i].Dist = idDist[strconv.FormatUint(shops[i].ID, 10)]
	}
	return dto.OkData(shops)
}

// ById The detailed information:
// @Title ById
// @Description shop根据id进行查询
// @Param id
// @Return tb.TbShop
func (Shop) ById(id int) dto.Response {
	key := global.ShopCache + strconv.Itoa(id)
	// 布隆过滤器判断是否有改key，解决缓存穿透
	if !global.Bloomfilter.Contain([]byte(key)) {
		return dto.Err(key + "不存在")
	}
	var shop tb.TbShop
	// 从redis缓存中获取店铺信息
	boolRes, _ := utils.EnterUtilsApp.RedisCacheUtils.GetCacheData(key, &shop)
	// redis锁的key
	lockValue := uuid.NewV4().String() + strconv.Itoa(id)
	// false说明没有该缓存
	if !boolRes {
		rlock := redisLock.NewRedisLock(key, lockValue, time.Second*2)
		// 获取锁，拿到锁的去数据库查询重建缓存，拿不到的等待
		for !rlock.TryLock() {
			time.Sleep(time.Millisecond * 50)
		}
		// DCL双重检查缓存。获取到锁再次检查是否有缓存，防止拿到锁后多次查询
		boolRes, _ := utils.EnterUtilsApp.RedisCacheUtils.GetCacheData(key, &shop)
		// 如果有缓存直接返回
		if boolRes {
			rlock.UnLock()
			return dto.OkData(shop)
		}
		// 模拟重建缓存需要很长时间
		//time.Sleep(time.Millisecond * 200)

		// 查询商铺信息
		affected := global.MysqlDb.Model(&tb.TbShop{}).Scopes(EnterServicesApp.PaginateService.byId(id)).Find(&shop).RowsAffected
		// 判断是否有该商铺
		if affected <= 0 {
			rlock.UnLock()
			return dto.Err("no shop id")
		}
		// json序列化shop存入redis
		utils.EnterUtilsApp.RedisCacheUtils.AddCacheDataAndSetTTL(key, shop)
		rlock.UnLock()
		return dto.OkData(shop)
	}
	return dto.OkData(shop)
}

// Save The detailed information:
// @Title Save
// @Description 添加商铺信息
// @Param shop
// @Return bool
func (Shop) Save(shop tb.TbShop) dto.Response {
	global.Bloomfilter.Add([]byte(global.ShopCache + strconv.Itoa(int(shop.ID))))
	tx := global.MysqlDb.Save(&shop)
	if tx.Error != nil {
		return dto.Err("insertShopErr")
	}
	return dto.Ok()
}

// Update The detailed information:
// @Title Update
// @Description 修改商铺信息
func (Shop) Update(shop tb.TbShop) dto.Response {

	return dto.Ok()
}

func (Shop) ShopOfName(name, current string) dto.Response {
	page, err := strconv.Atoi(current)
	if err != nil {
		page = 1
	}
	var shops []tb.TbShop
	global.MysqlDb.Model(&tb.TbShop{}).Scopes(EnterServicesApp.PaginateService.paging(page)).Find(&shops)
	return dto.OkData(shops)
}
