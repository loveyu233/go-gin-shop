package tb

type TbVoucherOrder struct {
	ID         int64  `json:"id" gorm:"column:id"`                   // 主键
	UserID     uint64 `json:"user_id" gorm:"column:user_id"`         // 下单的用户id
	VoucherID  uint64 `json:"voucher_id" gorm:"column:voucher_id"`   // 购买的代金券id
	PayType    uint8  `json:"pay_type" gorm:"column:pay_type"`       // 支付方式 1：余额支付；2：支付宝；3：微信
	Status     uint8  `json:"status" gorm:"column:status"`           // 订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款
	CreateTime string `json:"create_time" gorm:"column:create_time"` // 下单时间
	PayTime    string `json:"pay_time" gorm:"column:pay_time"`       // 支付时间
	UseTime    string `json:"use_time" gorm:"column:use_time"`       // 核销时间
	RefundTime string `json:"refund_time" gorm:"column:refund_time"` // 退款时间
	UpdateTime string `json:"update_time" gorm:"column:update_time"` // 更新时间
}

func (m *TbVoucherOrder) TableName() string {
	return "tb_voucher_order"
}
