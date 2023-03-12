package service

import (
	"go-gin-shop/global"
	"gorm.io/gorm"
)

// Paginate 公共模版
type Paginate struct {
}

// paging The detailed information:
// @Title paging
// @Description 分页查询模版
// @Param page 页码
// @Return func(db *gorm.DB) *gorm.DB
func (Paginate) paging(page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * global.NaxPageSize
		return db.Offset(offset).Limit(global.NaxPageSize)
	}
}

// byId The detailed information:
// @Title byId
// @Description 通过id进行查询的模版
// @Param id
// @Return func(db *gorm.DB) *gorm.DB
func (Paginate) byId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// byTypeId The detailed information:
// @Title byTypeId
// @Description 通过typeId进行查询的模版
// @Param typeId
// @Return func(db *gorm.DB) *gorm.DB
func (Paginate) byTypeId(typeId int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type_id = ?", typeId)
	}
}
