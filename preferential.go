package preferential

import (
	"github.com/shopspring/decimal"
)

const (
	PreferentialTypeNormal     = iota // 0-正常逻辑
	PreferentialTypeOnSale            // 1-打折
	PreferentialTypeOnceRebate        // 2-返利(满x返y,只返一次)
	PreferentialTypeLoopRebate        // 3-返利(每满x返y,返多次)
)

type SuperPreferential interface {
	// 计算最终订单数据
	// 参数 产品id,原始金额
	// 返回 应收金额,返利金额,错误
	Calculate(int64, decimal.Decimal) (decimal.Decimal, decimal.Decimal, error)
}

// 正常逻辑
type Normal struct{}

func (n Normal) Calculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	return src, decimal.Zero, nil
}

// 打折
type OnSale struct{}

func (os OnSale) Calculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	return onSaleCalculate(productID, src)
}

// 返利(满x返y,只返一次)
type OnceRebateCash struct{}

func (or OnceRebateCash) Calculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	return onceRebateCalculate(productID, src)
}

// 返利(每满x返y,返多次)
type LoopRebateCash struct{}

func (lr LoopRebateCash) Calculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	return loopRebateCalculate(productID, src)
}

// 工厂类
type OrderFactory struct {
	engine SuperPreferential
}

func (of *OrderFactory) Calculate(productID int64, src decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	return of.engine.Calculate(productID, src)
}

// 实例化订单工厂
func NewOrderFactory(preferentialType int) *OrderFactory {
	f := new(OrderFactory)
	switch preferentialType {
	case PreferentialTypeOnSale:
		f.engine = OnSale{}
	case PreferentialTypeOnceRebate:
		f.engine = OnceRebateCash{}
	case PreferentialTypeLoopRebate:
		f.engine = LoopRebateCash{}
	default:
		f.engine = Normal{}
	}
	return f
}
