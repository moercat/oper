package op

import (
	"github.com/spf13/cast"
)

func In(fat, sub interface{}) bool {
	_, ok := InI(fat, sub)
	return ok
}

func Index(fat, idx interface{}) int {
	index, _ := InI(fat, idx)
	return index
}

func Extend(fat, sub interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	subs := cast.ToStringSlice(sub)

	fats = append(fats, subs...)

	return fats
}

func Pop(fat interface{}, idx int) interface{} {
	fats := cast.ToStringSlice(fat)
	if idx < 0 {
		idx = len(fats) + idx
	}

	return append(fats[:idx], fats[(idx+1):]...)
}

func Remove(fat interface{}, value interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)
	for i, v := range fats {
		if v == str {
			fat = append(fats[:i], fats[(i+1):]...)
		}

	}

	return fat
}

func Append(fat interface{}, value interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)

	return append(fats, str)
}

func Insert(fat interface{}, idx int, value interface{}) interface{} {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)

	res := append(fats[:idx], str)
	fat = append(res, fats[idx:]...)

	return append(res, fats[idx:]...)
}

func Count(fat interface{}, value interface{}) (count int) {
	fats := cast.ToStringSlice(fat)
	str := cast.ToString(value)

	for _, v := range fats {
		if v == str {
			count += 1
		}
	}

	return count
}

func InI(fat interface{}, sub interface{}) (int, bool) {
	fats := cast.ToStringSlice(fat)
	s := cast.ToString(sub)

	for i, v := range fats {
		if v == s {
			return i, true
		}
	}

	return -1, false
}

func Equal(fat, sub interface{}) bool {
	s1 := cast.ToStringSlice(fat)
	s2 := cast.ToStringSlice(sub)
	if len(s1) != len(s2) {
		return false
	}
	for i, n := range s1 {
		if n != s2[i] {
			return false
		}
	}
	return true
}
