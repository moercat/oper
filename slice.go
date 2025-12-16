package oper

import "github.com/spf13/cast"

type Number interface {
	~int8 | ~int16 | ~int | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64
}

type List interface {
	Number | ~string
}

func In[N List](fat []N, sub N) bool {
	_, ok := InI(fat, sub)
	return ok
}

func Index[N List](fat []N, idx N) int {
	index, _ := InI(fat, idx)
	return index
}

func Extend[N List](fat []N, sub []N) []N {

	fat = append(fat, sub...)

	return fat
}

func Pop[N List](fat []N, idx int) []N {
	if idx >= len(fat) || len(fat)+idx < 0 {
		return nil
	}

	if idx < 0 {
		idx = len(fat) + idx
	}

	return append(fat[:idx], fat[(idx+1):]...)
}

func Remove[N List](fat []N, value N) []N {
	for i, v := range fat {
		if v == value {
			fat = append(fat[:i], fat[(i+1):]...)
		}
	}

	return fat
}

func Append[N List](fat []N, value N) []N {

	return append(fat, value)
}

func Insert[N List](fat []N, idx int, value N) []N {
	// 检查索引是否有效
	if idx < 0 {
		idx = len(fat) + idx + 1 // 负数索引从末尾开始计算
		if idx < 0 {
			idx = 0 // 如果计算后的索引仍小于0，则插入到开头
		}
	}
	if idx > len(fat) {
		idx = len(fat) // 如果索引超过长度，则插入到末尾
	}

	// 分配足够容量避免多次分配
	result := make([]N, len(fat)+1)

	// 复制前面部分
	copy(result, fat[:idx])

	// 插入新值
	result[idx] = value

	// 复制后面部分
	copy(result[idx+1:], fat[idx:])

	return result
}

func Count[N List](fat []N, value N) (count int) {

	for _, v := range fat {
		if v == value {
			count += 1
		}
	}

	return count
}

func InI[N List](fat []N, sub N) (int, bool) {

	for i, v := range fat {
		if v == sub {
			return i, true
		}
	}

	return -1, false
}

func Equal[N List](fat []N, sub []N) bool {
	if len(fat) != len(sub) {
		return false
	}
	for i, n := range fat {
		if n != sub[i] {
			return false
		}
	}
	return true
}

// Map 对切片中的每个元素应用函数并返回新切片 - 类似Python和JavaScript的map
func Map[T, R List](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter 过滤满足条件的元素 - 类似Python和JavaScript的filter
func Filter[T List](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Any 检查是否有任一元素满足条件 - 类似Python的any
func Any[T List](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// All 检查是否所有元素都满足条件 - 类似Python的all
func All[T List](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Find 查找满足条件的第一个元素及其索引 - 类似JavaScript的find
func Find[T List](slice []T, fn func(T) bool) (T, int) {
	var zero T
	for i, v := range slice {
		if fn(v) {
			return v, i
		}
	}
	return zero, -1
}

// FindIndex 查找满足条件的第一个元素的索引 - 类似JavaScript的findIndex
func FindIndex[T List](slice []T, fn func(T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

// Some 检查是否有任一元素满足条件 - 类似JavaScript的some
func Some[T List](slice []T, fn func(T) bool) bool {
	return Any(slice, fn)
}

// Every 检查是否所有元素都满足条件 - 类似JavaScript的every
func Every[T List](slice []T, fn func(T) bool) bool {
	return All(slice, fn)
}

// Chunk 将切片分块 - 类似Lodash的chunk
func Chunk[T List](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}

	var chunks [][]T
	for size < len(slice) {
		slice, chunks = slice[size:], append(chunks, slice[0:size:size])
	}

	if len(slice) > 0 {
		chunks = append(chunks, slice)
	}

	return chunks
}

// Flatten 将二维切片扁平化为一维 - 类似Lodash的flatten
func Flatten[T List](slice [][]T) []T {
	var result []T
	for _, s := range slice {
		result = append(result, s...)
	}
	return result
}

// UniqBy 基于提供的比较函数去重 - 类似Lodash的uniqBy
func UniqBy[T List](slice []T, fn func(T) T) []T {
	var result []T
	seen := make(map[string]bool)

	for _, v := range slice {
		key := fn(v) // 应用函数获得比较键
		keyStr := cast.ToString(key)
		if !seen[keyStr] {
			seen[keyStr] = true
			result = append(result, v)
		}
	}

	return result
}

// GroupBy 按照指定条件分组 - 类似Lodash的groupBy
func GroupBy[T List, K comparable](slice []T, fn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Contains 检查切片是否包含指定元素 - 类似各种语言的includes方法
func Contains[T List](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// Reverse 反转切片元素的顺序
func Reverse[T List](slice []T) []T {
	length := len(slice)
	result := make([]T, length)

	for i, v := range slice {
		result[length-i-1] = v
	}

	return result
}

// Take 获取切片的前n个元素
func Take[T List](slice []T, n int) []T {
	if n > len(slice) {
		n = len(slice)
	}
	if n < 0 {
		n = 0
	}
	return slice[:n]
}

// Drop 丢弃切片的前n个元素
func Drop[T List](slice []T, n int) []T {
	if n > len(slice) {
		return []T{}
	}
	if n < 0 {
		n = 0
	}
	return slice[n:]
}
