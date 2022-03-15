package oper

import "github.com/spf13/cast"

func Set(fat interface{}) (sub interface{}) {
	var subs []string
	fats := cast.ToStringSlice(fat)

	maps := make(map[string]bool)
	for _, v := range fats {
		if !maps[v] {
			maps[v] = true
			subs = append(subs, v)
		}
	}
	return subs
}

func Add(fat interface{}, value interface{}) (sub interface{}) {
	var subs []string
	fats := cast.ToStringSlice(fat)
	va := cast.ToString(value)

	maps := make(map[string]bool)
	fats = append(fats, va)
	for _, v := range fats {
		if !maps[v] {
			maps[v] = true
			subs = append(subs, v)
		}
	}
	return subs
}

func Union(fat interface{}, set ...interface{}) (sub interface{}) {
	var subs []string
	fats := cast.ToStringSlice(fat)
	for _, v := range set {
		sets := cast.ToStringSlice(v)
		fats = append(fats, sets...)
	}

	maps := make(map[string]bool)
	for _, v := range fats {
		if !maps[v] {
			maps[v] = true
			subs = append(subs, v)
		}
	}
	return subs
}

func Update(fat interface{}, set interface{}) (sub interface{}) {
	var subs []string
	fats := cast.ToStringSlice(fat)
	sets := cast.ToStringSlice(set)
	fats = append(fats, sets...)

	maps := make(map[string]bool)
	for _, v := range fats {
		if !maps[v] {
			maps[v] = true
			subs = append(subs, v)
		}
	}
	return subs
}
