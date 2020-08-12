package preferential

import (
	"github.com/shopspring/decimal"
)

// 折扣计算主逻辑
func onSaleCalculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	var real, rebate decimal.Decimal

	// 根据产品id获取对应的折扣规则
	ratio, err := getOnSaleRule(productID)
	if err != nil {
		return real, rebate, err
	}
	if ratio.IsZero() {
		// 如果得到的ratio为0,那么设置成1,防止出现白给
		ratio = decimal.New(1,0)
	}
	return src.Mul(ratio), decimal.Zero, nil
}

func getOnSaleRule(productID int64) (decimal.Decimal, error) {
	// 这里改为从缓存或数据库获取最终结果
	return decimal.RequireFromString("0.95"), nil
}