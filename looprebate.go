package preferential

import (
	"github.com/shopspring/decimal"
)

// 返利计算主逻辑
func loopRebateCalculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	// 根据产品id获取对应的返利规则
	term, reward, err := getLoopRebateRule(productID)
	if err != nil {
		return src, decimal.Zero, err
	}

	if reward.LessThanOrEqual(decimal.Zero) {
		// 如果满X返Y,Y小于等于0,那么直接返回
		return src, decimal.Zero, nil
	}

	if term.LessThanOrEqual(decimal.Zero) {
		// 如果满X返Y,X小于0,那么可能是下单即返
		return src, reward, nil
	}

	if src.LessThan(term) {
		// 不满足条件, 直接返回
		return src, decimal.Zero, nil
	}

	return src, reward.Mul(src.Div(term).Truncate(0)), nil
}

func getLoopRebateRule(productID int64) (decimal.Decimal, decimal.Decimal, error) {
	// 这里改为从缓存或数据库获取最终结果
	return decimal.RequireFromString("10"), decimal.RequireFromString("1"), nil
}