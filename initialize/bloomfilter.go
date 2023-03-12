package initialize

import (
	"github.com/linvon/cuckoo-filter"
	"go-gin-shop/enter/tb"
	"go-gin-shop/global"

	"strconv"
)

func InitBloomfilter() *cuckoo.Filter {
	var shops []tb.TbShop
	global.MysqlDb.Model(&tb.TbShop{}).Find(&shops)
	var voucher []tb.TbVoucher
	global.MysqlDb.Model(&tb.TbVoucher{}).Find(&voucher)
	cf := cuckoo.NewFilter(4, 9, uint((len(shops)+len(voucher))*3), cuckoo.TableTypePacked)
	for _, v := range shops {
		k := global.ShopCache + strconv.Itoa(int(v.ID))
		cf.Add([]byte(k))
	}
	for _, v := range voucher {
		k := global.VoucherCache + strconv.Itoa(int(v.ID))
		cf.Add([]byte(k))
	}
	return cf
}
