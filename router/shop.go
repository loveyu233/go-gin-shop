package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
)

// ShopRouter shop路由集中注册
type ShopRouter struct{}

// InitShopRouter The detailed information:
// @Title InitShopRouter
// @Description 初始化shop路由
// @Param Router
func (ShopRouter) InitShopRouter(Router *gin.RouterGroup) {
	shopType := Router.Group("shop-type")
	showApi := api.ApiGroupApp.ShopApi
	{
		// 商品分类信息
		shopType.GET("/list", showApi.ShowType)
	}
	shop := Router.Group("/shop")
	{
		// 查询某一类的商铺信息
		shop.GET("/of/type", showApi.OfType)
		// 查询具体商铺信息
		shop.GET("/:id", showApi.ById)
		// 保存商铺信息
		shop.POST("", showApi.SaveShop)
		// 修改商铺信息
		shop.PUT("", showApi.UpdateShop)
		// 查询和我相关的商铺信息
		shop.GET("/of/name", showApi.ShopOfName)
	}
}
