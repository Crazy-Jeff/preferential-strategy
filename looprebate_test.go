package preferential

import (
	"testing"

	"bou.ke/monkey"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestLoopRebateCash_Calculate(t *testing.T) {
	engine := NewOrderFactory(PreferentialTypeLoopRebate)
	Convey("每满X返Y优惠", t, func() {
		Convey("查询产品优惠规则失败", func() {
			monkey.Patch(getLoopRebateRule, func(productID int64) (decimal.Decimal, decimal.Decimal, error) {
				return decimal.Zero, decimal.Zero, errors.New("not fund")
			})
			src := decimal.RequireFromString("123")
			_, _, err := engine.Calculate(0, src)
			assert.NotNil(t, err, "正常返回错误")
		})
		Convey("每满X返Y优惠, Y为0", func() {
			monkey.Patch(getLoopRebateRule, func(productID int64) (decimal.Decimal, decimal.Decimal, error) {
				return decimal.Zero, decimal.Zero, nil
			})
			src := decimal.RequireFromString("123")
			real, rebate, err := engine.Calculate(0, src)
			assert.Nil(t, err, "应无报错")
			assert.Equal(t, src, real, "实付数额应相等")
			assert.Equal(t, rebate, decimal.Zero, "返利应为0")
		})
		Convey("每满X返Y优惠, X为0, Y为10", func() {
			monkey.Patch(getLoopRebateRule, func(productID int64) (decimal.Decimal, decimal.Decimal, error) {
				return decimal.Zero, decimal.RequireFromString("10"), nil
			})
			src := decimal.RequireFromString("123")
			real, rebate, err := engine.Calculate(0, src)
			assert.Nil(t, err, "应无报错")
			assert.Equal(t, src, real, "实付数额应相等")
			assert.Equal(t, rebate, decimal.RequireFromString("10"), "返利应为10")
		})
		Convey("每满X返Y优惠, X为10, Y为1, 传入金额小于X", func() {
			monkey.Patch(getLoopRebateRule, func(productID int64) (decimal.Decimal, decimal.Decimal, error) {
				return decimal.RequireFromString("10"), decimal.RequireFromString("1"), nil
			})
			src := decimal.RequireFromString("9.999")
			real, rebate, err := engine.Calculate(0, src)
			assert.Nil(t, err, "应无报错")
			assert.Equal(t, src, real, "实付数额应相等")
			assert.Equal(t, rebate, decimal.Zero, "应无返利")
		})
		Convey("每满X返Y优惠, X为10, Y为1", func() {
			monkey.Patch(getLoopRebateRule, func(productID int64) (decimal.Decimal, decimal.Decimal, error) {
				return decimal.RequireFromString("10"), decimal.RequireFromString("1"), nil
			})
			src := decimal.RequireFromString("129")
			real, rebate, err := engine.Calculate(0, src)
			assert.Nil(t, err, "应无报错")
			assert.Equal(t, src, real, "实付数额应相等")
			assert.Equal(t, rebate, decimal.RequireFromString("12"), "返利应为12(129/10*1)")
		})
	})
}
