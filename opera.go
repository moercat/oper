package oper

import (
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// Opera item + 为 乘 - 为除
func Opera[N List](str N, item int64) string {
	st := cast.ToString(str)
	deci, _ := decimal.NewFromString(st)
	i := decimal.NewFromInt(item)

	if item < 0 {
		i = decimal.NewFromInt(-item)
		return deci.Div(i).String()
	}

	return deci.Mul(i).String()
}
