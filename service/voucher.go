package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-gin-shop/enter/dto"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"
	"go-gin-shop/utils"
	"gorm.io/gorm"

	"strconv"
)

var flag = false

// VouCher 优惠卷
type VouCher struct {
}

// List The detailed information:
// @Title List
// @Description 根据id查询优惠卷
// @Param id
// @Return []tb.TbVoucher
func (v VouCher) List(id int) dto.Response {
	key := global.VoucherCache + strconv.Itoa(id)
	if !global.Bloomfilter.Contain([]byte(key)) {
		return dto.Err(key + "不存在")
	}
	tx := global.MysqlDb.Begin()
	var vouchers []tb.TbVoucher
	res := tx.Model(&tb.TbVoucher{}).Where("shop_id = ?", id).Find(&vouchers)
	if res.Error != nil {
		tx.Rollback()
		return dto.Err("数据查询失败")
	}
	if len(vouchers) >= 1 {
		for i := range vouchers {
			seckillVoucher := tb.TbSeckillVoucher{}
			res = tx.Model(&tb.TbSeckillVoucher{}).Where("voucher_id = ?", vouchers[i].ID).First(&seckillVoucher)
			if res.RowsAffected >= 1 {
				vouchers[i].Stock = seckillVoucher.Stock
				vouchers[i].BeginTime = seckillVoucher.BeginTime
				vouchers[i].EndTime = seckillVoucher.EndTime
			}
			if res.Error != nil {
				tx.Rollback()
				return dto.Err("数据查询失败")
			}
		}
	}
	if tx.Commit().Error != nil {
		return dto.Err("数据查询失败")
	}
	return dto.OkData(vouchers)
}

// AddVoucher The detailed information:
// @Title Seckill
// @Description 添加优惠卷
// @Param voucher
// @Return dto.Response
func (VouCher) AddVoucher(voucher tb.TbVoucher) dto.Response {
	// 添加到过滤器中
	global.Bloomfilter.Add([]byte(global.VoucherCache + strconv.Itoa(int(voucher.ID))))
	// 获取当前时间
	timeNow := global.TimeNow
	// 1: 上架中
	voucher.Status = 1
	// 创建时间
	voucher.CreateTime = timeNow
	// 开启事务
	tx := global.MysqlDb.Begin()
	if tx.Create(&voucher).Error != nil {
		// 回滚事务
		tx.Rollback()
	}
	// 创建秒杀优惠卷数据
	ev := tb.TbSeckillVoucher{
		VoucherID:  voucher.ID,         //优惠卷id
		Stock:      voucher.Stock,      //库存
		CreateTime: voucher.CreateTime, //创建时间
		BeginTime:  voucher.BeginTime,  //开始时间
		EndTime:    voucher.EndTime,    //结束时间
	}
	if tx.Create(&ev).Error != nil {
		tx.Rollback()
	}
	// 提交事务
	if tx.Commit().Error != nil {
		return dto.Err("err")
	}
	// 优惠卷信息添加到redis中
	global.RedisDb.Set(global.Content, global.VoucherRedisKey+strconv.FormatUint(voucher.ID, 10), voucher.Stock, 0)
	return dto.OkData(voucher.ID)
}

// VoucherOrderId The detailed information:
// @Title VoucherOrderId
// @Description 抢购优惠卷
// @Param id 优惠卷id
// @Param c 获取jwt来拿到用户信息
// @Return dto.Response 返回订单id
func (VouCher) VoucherOrderId(voucherId string, c *gin.Context) dto.Response {
	// 获取用户信息
	user := utils.EnterUtilsApp.JwtUtils.JwtGetUser(c)
	// 执行lua脚本判断用户是否有资格抢购优惠卷; 返回值： 0：有资格；1：没有库存；2：已经抢购过了；
	result, _ := redis.NewScript(global.IsQualificationLuaScript).Run(global.Content, global.RedisDb, []string{}, voucherId, user.ID).Result()
	if result == nil {
		return dto.Err("不存在该优惠卷")
	}
	// 根据返回值来判断
	switch result.(int64) {
	case 1:
		return dto.Err("库存不足")
	case 2:
		return dto.Err("一人只可抢购一单")
	}
	// 封装订单信息
	voucherOrder := returnVoucherOrder(user.ID, voucherId)
	err := createVoucherOrder(voucherOrder)
	if err != nil {
		return dto.Err("订单保存失败")
	}
	// 返回订单ID
	return dto.OkData(voucherOrder.ID)
}

// returnVoucherOrder The detailed information:
// @Title returnVoucherOrder
// @Description 根据用户id和优惠卷id创建订单信息
// @Param userId
// @Param voucherID
// @Return *tb.TbVoucherOrder
func returnVoucherOrder(userId uint64, voucherID string) *tb.TbVoucherOrder {
	// 新增订单
	voucherOrder := tb.TbVoucherOrder{}
	// 订单ID，全局唯一ID
	voucherOrder.ID = utils.EnterUtilsApp.RedisIDWorker.NextId("voucherOrder")
	// 用户ID
	voucherOrder.UserID = userId
	// 优惠卷ID
	id, _ := strconv.Atoi(voucherID)
	voucherOrder.VoucherID = uint64(id)
	// 默认为 未支付、余额支付、订单创建时间
	voucherOrder.Status = 1
	voucherOrder.PayType = 1
	voucherOrder.CreateTime = global.TimeNow
	// 返回订单结构体
	return &voucherOrder
}

// createVoucherOrder The detailed information:
// @Title createVoucherOrder
// @Description 数据库创建订单
// @Param voucherOrder
// @Return error
func createVoucherOrder(voucherOrder *tb.TbVoucherOrder) error {
	// 开启事务
	tx := global.MysqlDb.Begin()
	defer func() {
		// 捕捉错误
		err := recover()
		if err != nil {
			// 回滚
			tx.Rollback()
		}
	}()
	// 修改订单，这里用 stock > 0 来判断，用数据库的锁机制来实现乐观锁解决超卖问题
	resRowsAffected := tx.Model(&tb.TbSeckillVoucher{}).Where("voucher_id = ? and stock > 0", voucherOrder.VoucherID).Update("stock", gorm.Expr("stock - 1")).RowsAffected
	// 影响行数
	if resRowsAffected == 0 {
		// 回滚
		tx.Rollback()
		return errors.New("修改订单失败")
	}
	// 保存订单到数据库
	resRowsAffected = tx.Model(&tb.TbVoucherOrder{}).Create(&voucherOrder).RowsAffected
	// 影响行数
	if resRowsAffected == 0 {
		// 回滚
		tx.Rollback()
		return errors.New("保存订单失败")
	}
	// 事务提交
	tx.Commit()
	// 正常返回为nil
	return nil
}
