package oper

import (
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// Opera item > 0 为 乘法, item < 0 为除法
func Opera[N List](str N, item int64) string {
	st := cast.ToString(str)
	deci, err := decimal.NewFromString(st)
	if err != nil {
		// 如果不能转换为decimal，则返回原字符串
		return st
	}

	i := decimal.NewFromInt(item)

	if item < 0 {
		i = decimal.NewFromInt(-item)
		return deci.Div(i).String()
	}

	return deci.Mul(i).String()
}
