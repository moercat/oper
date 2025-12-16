package oper

import "github.com/spf13/cast"

// Set 对切片进行去重操作，返回不包含重复元素的切片
func Set(fat interface{}) []string {
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

// Add 向集合中添加一个元素，如果元素已存在则不重复添加
func Add(fat interface{}, value interface{}) []string {
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

// Union 将多个集合合并成一个去重后的集合
func Union(fat interface{}, set ...interface{}) []string {
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

// Update 更新集合，将另一个集合并入当前集合（去重）
func Update(fat interface{}, set interface{}) []string {
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
