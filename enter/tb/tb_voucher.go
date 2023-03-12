package tb

type TbVoucher struct {
	ID          uint64 `json:"id" gorm:"column:id"`                    // 主键
	ShopID      uint64 `json:"shopId" gorm:"column:shop_id"`           // 商铺id
	Title       string `json:"title" gorm:"column:title"`              // 代金券标题
	SubTitle    string `json:"subTitle" gorm:"column:sub_title"`       // 副标题
	Rules       string `json:"rules" gorm:"column:rules"`              // 使用规则
	PayValue    uint64 `json:"payValue" gorm:"column:pay_value"`       // 支付金额，单位是分。例如200代表2元
	ActualValue int64  `json:"actualValue" gorm:"column:actual_value"` // 抵扣金额，单位是分。例如200代表2元
	Type        uint8  `json:"type" gorm:"column:type"`                // 0,普通券；1,秒杀券
	Status      uint8  `json:"status" gorm:"column:status"`            // 1,上架; 2,下架; 3,过期
	CreateTime  string `json:"createTime" gorm:"column:create_time"`   // 创建时间
	UpdateTime  string `json:"updateTime" gorm:"column:update_time"`   // 更新时间
	Stock       int    `json:"stock" gorm:"-"`                         // 库存
	BeginTime   string `json:"beginTime"  gorm:"-"`                    // 生效时间
	EndTime     string `json:"endTime"  gorm:"-"`                      // 失效时间
}

func (m *TbVoucher) TableName() string {
	return "tb_voucher"
}
