package oper

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// Between 数值是否在某个范围内,将自动转换类型再比较.
func Between(value interface{}, left interface{}, right interface{}) (bool, bool, error) {

	v, _ := decimal.NewFromString(cast.ToString(value))
	l, _ := decimal.NewFromString(cast.ToString(left))
	r, _ := decimal.NewFromString(cast.ToString(right))

	if v.LessThanOrEqual(l) || v.GreaterThanOrEqual(r) {
		return true, false, errors.New("数值需大于" + l.String() + "小于" + r.String())
	}

	if v.LessThan(l) || v.GreaterThan(r) {
		return false, true, errors.New("数值需大于等于" + l.String() + "小于等于" + r.String())
	}

	return false, false, nil
}
