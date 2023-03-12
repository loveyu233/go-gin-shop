package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/tb"
	"go-gin-shop/service"

	"strconv"
)

type VoucherApi struct {
}

// List The detailed information:
// @Title List
// @Description 通过商铺id返回对应的优惠卷
// @Param c
func (v VoucherApi) List(c *gin.Context) {
	// 商品id
	shopId, _ := strconv.Atoi(c.Param("shopId"))
	c.JSON(200, service.EnterServicesApp.VouCherService.List(shopId))
}

// Seckill The detailed information:
// @Title Seckill
// @Description 添加优惠卷包括秒杀卷
// @Param c
func (VoucherApi) Seckill(c *gin.Context) {
	voucher := tb.TbVoucher{}
	err := c.ShouldBindJSON(&voucher)
	if err != nil {
		c.JSON(200, dto.Err(err.Error()))
		return
	}
	c.JSON(200, service.EnterServicesApp.VouCherService.AddVoucher(voucher))
}

func (VoucherApi) VoucherOrderId(c *gin.Context) {
	voucherId := c.Param("id")
	c.JSON(200, service.EnterServicesApp.VouCherService.VoucherOrderId(voucherId, c))
}
