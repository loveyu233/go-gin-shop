package tb

type TbSeckillVoucher struct {
	VoucherID  uint64 `json:"voucher_id" gorm:"column:voucher_id"`   // 关联的优惠券的id
	Stock      int    `json:"stock" gorm:"column:stock"`             // 库存
	CreateTime string `json:"create_time" gorm:"column:create_time"` // 创建时间
	BeginTime  string `json:"begin_time" gorm:"column:begin_time"`   // 生效时间
	EndTime    string `json:"end_time" gorm:"column:end_time"`       // 失效时间
	UpdateTime string `json:"update_time" gorm:"column:update_time"` // 更新时间
}

func (m *TbSeckillVoucher) TableName() string {
	return "tb_seckill_voucher"
}
