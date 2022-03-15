package list

import (
	"github.com/spf13/cast"
)

type List []interface{}

func (l *List) Create(fat interface{}) {
	*l = append(*l, fat)
}

func (l *List) In(sub interface{}) bool {
	tmp := *l
	_, ok := InI(tmp[0], sub)
	return ok
}

func (l *List) Index(sub interface{}) int {
	tmp := *l
	index, _ := InI(tmp[0], sub)
	return index
}

func (l *List) Extend(sub interface{}) interface{} {
	tmp := *l
	fats := cast.ToStringSlice(tmp[0])
	subs := cast.ToStringSlice(sub)

	fats = append(fats, subs...)

	return fats
}

func (l *List) Pop(idx int) interface{} {
	tmp := *l
	fats := cast.ToStringSlice(tmp[0])
	if idx < 0 {
		idx = len(fats) + idx
	}
	tmp[0] = append(fats[:idx], fats[(idx+1):]...)
	l = &tmp
	return l
}

func (l *List) Remove(value interface{}) interface{} {
	tmp := *l
	fats := cast.ToStringSlice(tmp[0])
	str := cast.ToString(value)
	for i, v := range fats {
		if v == str {
			tmp[0] = append(fats[:i], fats[(i+1):]...)
		}

	}

	l = &tmp
	return l
}

func (l *List) Append(value interface{}) interface{} {
	tmp := *l
	fats := cast.ToStringSlice(tmp[0])
	str := cast.ToString(value)

	tmp[0] = append(fats, str)
	l = &tmp
	return append(fats, str)
}

func (l *List) Insert(idx int, value interface{}) interface{} {
	tmp := *l
	fats := cast.ToStringSlice(tmp[0])
	str := cast.ToString(value)

	res := append(fats[:idx], str)
	tmp[0] = append(res, fats[idx:]...)

	l = &tmp
	return tmp[0]
}

func (l *List) Count(value interface{}) (count int) {
	tmp := *l
	fats := cast.ToStringSlice(tmp[0])
	str := cast.ToString(value)

	for _, v := range fats {
		if v == str {
			count += 1
		}
	}

	return count
}

func InI(fat interface{}, sub interface{}) (int, bool) {
	idx := cast.ToString(sub)
	fats := cast.ToStringSlice(fat)

	for i, s := range fats {
		if idx == s {
			return i, true
		}
	}

	return -1, false
}
