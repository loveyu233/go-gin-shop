package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/api"
	"go-gin-shop/middleware"
)

// VoucherRouter 优惠卷路由集中注册
type VoucherRouter struct {
}

// InitVoucherRouter The detailed information:
// @Title InitVoucherRouter
// @Description 初始化优惠卷路由
// @Param Group
func (v VoucherRouter) InitVoucherRouter(Group *gin.RouterGroup) {
	voucher := Group.Group("/voucher")
	voucher.GET("/list/:shopId", api.ApiGroupApp.VoucherApi.List)
	voucher.POST("/seckill", api.ApiGroupApp.VoucherApi.Seckill)

	voucherOrder := Group.Group("/voucher-order")
	voucherOrder.Use(middleware.Auth)
	voucherOrder.POST("/seckill/:id", api.ApiGroupApp.VoucherApi.VoucherOrderId)

}
