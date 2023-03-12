package utils

import (
	"encoding/json"
	"go-gin-shop/enter/tb"
)

type StructMapUtils struct {
}

func (StructMapUtils) StructToMap(user tb.TbUser) (m map[string]interface{}) {
	marshal, _ := json.Marshal(user)
	json.Unmarshal(marshal, &m)
	return m
}

func (StructMapUtils) MapToStruct(userMap *map[string]string) (user *tb.TbUser) {
	marshal, err := json.Marshal(userMap)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(marshal, &user)
	if err != nil {
		return nil
	}
	return
}
