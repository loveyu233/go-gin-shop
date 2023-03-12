package utils

import (
	"encoding/json"
	"go-gin-shop/global"
)

type RedisCacheUtils struct {
}

func (RedisCacheUtils) AddCacheDataAndSetTTL(key string, structValue any) error {
	marshal, err := json.Marshal(structValue)
	if err != nil {
		return err
	}
	_, err = global.RedisDb.Set(global.Content, key, string(marshal), global.TTLTime).Result()
	if err != nil {
		return err
	}
	return nil
}

func (RedisCacheUtils) GetCacheData(key string, res any) (bool, error) {
	result, err := global.RedisDb.Get(global.Content, key).Result()
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(result), &res)
	if err != nil {
		return false, err
	}
	return true, nil
}
