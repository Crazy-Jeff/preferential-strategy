package preferential

import (
	"testing"

	"bou.ke/monkey"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestOnSale_Calculate(t *testing.T) {
	engine := NewOrderFactory(PreferentialTypeOnSale)
	Convey("打折优惠", t, func() {
		Convey("查询产品优惠规则失败", func() {
			monkey.Patch(getOnSaleRule, func(productID int64) (decimal.Decimal, error) {
				return decimal.Zero, errors.New("not fund")
			})
			src := decimal.RequireFromString("123")
			_, _, err := engine.Calculate(0, src)
			assert.NotNil(t, err, "正常返回错误")
		})
		Convey("打折 比例0", func() {
			monkey.Patch(getOnSaleRule, func(productID int64) (decimal.Decimal, error) {
				return decimal.Zero, nil
			})
			src := decimal.RequireFromString("1234")
			real, rebate, err := engine.Calculate(0, src)
			assert.Nil(t, err, "应无报错")
			assert.Equal(t, src, real, "实付数额应相等")
			assert.Equal(t, rebate, decimal.Zero, "返利应为0")
		})
		Convey("打折 比例0.9", func() {
			monkey.Patch(getOnSaleRule, func(productID int64) (decimal.Decimal, error) {
				return decimal.RequireFromString("0.9"), nil
			})
			src := decimal.RequireFromString("100")
			real, rebate, err := engine.Calculate(0, src)
			assert.Nil(t, err, "应无报错")
			assert.Equal(t, real, src.Mul(decimal.New(9, -1)), "实付数额应为90(100*0.9)")
			assert.Equal(t, rebate, decimal.Zero, "返利应为0")
		})
	})
}
