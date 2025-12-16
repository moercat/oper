package oper

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// Between 数值是否在某个范围内,将自动转换类型再比较.
// 返回值: (outLeftOrRight bool, outStrict bool, error)
// outLeftOrRight: 是否超出左边界或右边界
// outStrict: 是否严格超出边界(即 <= 或 >=)
func Between(value interface{}, left interface{}, right interface{}) (bool, bool, error) {
	v, err := decimal.NewFromString(cast.ToString(value))
	if err != nil {
		return true, false, errors.New("值不能转换为有效数字")
	}

	l, err := decimal.NewFromString(cast.ToString(left))
	if err != nil {
		return true, false, errors.New("左边界不能转换为有效数字")
	}

	r, err := decimal.NewFromString(cast.ToString(right))
	if err != nil {
		return true, false, errors.New("右边界不能转换为有效数字")
	}

	// 检查是否在开区间内 (left, right)
	if v.LessThanOrEqual(l) || v.GreaterThanOrEqual(r) {
		if v.LessThan(l) || v.GreaterThan(r) {
			// 严格超出边界
			return true, true, errors.New("数值需大于" + l.String() + "小于" + r.String())
		} else {
			// 等于边界
			return true, false, errors.New("数值需大于" + l.String() + "小于" + r.String())
		}
	}

	return false, false, nil
}
