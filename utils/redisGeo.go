package utils

import (
	"github.com/go-redis/redis/v8"
	"go-gin-shop/global"
)

type RedisGeoUtil struct {
}

// SetGeo The detailed information:
// @Title SetGeo
// @Description 添加
// @Param key
// @Param value
// @Param x
// @Param y
// @Return bool
func (RedisGeoUtil) SetGeo(key, value string, x, y float64) bool {
	result, err := global.RedisDb.GeoAdd(global.Content, key, &redis.GeoLocation{
		Name:      value,
		Longitude: x,
		Latitude:  y,
	}).Result()
	if err != nil {
		return false
	}
	if result <= 0 {
		return false
	}
	return true
}

func (RedisGeoUtil) SearchGeo(key string, x, y float64, current int) []redis.GeoLocation {
	from := (current - 1) * global.NaxPageSize
	end := current * global.NaxPageSize
	result, err := global.RedisDb.GeoSearchLocation(global.Content, key, &redis.GeoSearchLocationQuery{
		GeoSearchQuery: redis.GeoSearchQuery{
			Longitude:  x,
			Latitude:   y,
			Radius:     10000,
			RadiusUnit: "m",
			Count:      end,
		},
		WithCoord: false,
		WithDist:  true,
		WithHash:  false,
	}).Result()
	if err != nil || len(result) == 0 || from > len(result) {
		return []redis.GeoLocation{}
	}
	return result[from:]
}
