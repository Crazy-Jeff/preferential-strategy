package preferential

import (
	"github.com/shopspring/decimal"
)

// 返利计算主逻辑
func onceRebateCalculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	// 根据产品id获取对应的返利规则
	term, reward, err := getOnceRebateRule(productID)
	if err != nil {
		return src, decimal.Zero, err
	}

	if reward.LessThanOrEqual(decimal.Zero) {
		// 如果满X返Y,Y小于等于0,那么直接返回
		return src, decimal.Zero, nil
	}

	if src.LessThan(term) {
		// 不满足条件, 直接返回
		return src, decimal.Zero, nil
	}

	return src, reward, nil
}

func getOnceRebateRule(productID int64) (decimal.Decimal, decimal.Decimal, error) {
	// 这里改为从缓存或数据库获取最终结果
	return decimal.RequireFromString("10"), decimal.RequireFromString("1"), nil
}