package preferential

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNormal_Calculate(t *testing.T) {
	engine := NewOrderFactory(PreferentialTypeNormal)
	src := decimal.RequireFromString("123")
	real, rebate, err := engine.Calculate(0, src)
	assert.Nil(t, err, "应无报错")
	assert.Equal(t, src, real, "实付数额应相等")
	assert.Equal(t, rebate, decimal.Zero, "返利应为0")
}
