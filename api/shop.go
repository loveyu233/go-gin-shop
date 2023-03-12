package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/tb"
	"go-gin-shop/service"

	"strconv"
)

type ShopApi struct{}

// ShowType The detailed information:
// @Title ShowType
// @Description 商品分类信息
// @Param c
func (ShopApi) ShowType(c *gin.Context) {
	c.JSON(200, service.EnterServicesApp.ShopService.ShowType())
}

// OfType The detailed information:
// @Title OfType
// @Description 商品某一类型的数据
// @Param c
func (ShopApi) OfType(c *gin.Context) {
	// 商品类型
	id := c.Query("typeId")
	// 页码
	current := c.Query("current")

	x := c.Query("x")
	y := c.Query("y")

	c.JSON(200, service.EnterServicesApp.ShopService.OfType(id, current, x, y))
}

// ById The detailed information:
// @Title ById
// @Description 根据店铺id查询
// @Param c
func (ShopApi) ById(c *gin.Context) {
	// 店铺具体id
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, service.EnterServicesApp.ShopService.ById(id))
}

// SaveShop The detailed information:
// @Title SaveShop
// @Description 添加商铺信息
// @Param c
func (ShopApi) SaveShop(c *gin.Context) {
	// TODO
}

// UpdateShop The detailed information:
// @Title UpdateShop
// @Description 修改商铺信息
func (ShopApi) UpdateShop(c *gin.Context) {
	//TODO
	shop := tb.TbShop{}
	err := c.ShouldBindJSON(&shop)
	if err != nil {
		c.JSON(200, dto.Err("json err"))
	}
	c.JSON(200, service.EnterServicesApp.ShopService.Update(shop))
}

func (ShopApi) ShopOfName(c *gin.Context) {
	//TODO
	c.JSON(200, service.EnterServicesApp.ShopService.ShopOfName(c.Query("name"), c.Query("current")))
}
