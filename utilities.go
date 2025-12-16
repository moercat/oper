package oper

import (
	"regexp"
	"strings"
)

// String utilities

// Capitalize 首字母大写，其余字母小写
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// CamelCase 将字符串转换为驼峰格式
func CamelCase(s string) string {
	// 将非字母数字字符替换为空格
	spaceRegexp := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	parts := strings.Fields(spaceRegexp.ReplaceAllString(s, " "))

	if len(parts) == 0 {
		return ""
	}

	result := strings.ToLower(parts[0])
	for _, part := range parts[1:] {
		result += Capitalize(strings.ToLower(part))
	}

	return result
}

// KebabCase 将字符串转换为短横线格式
func KebabCase(s string) string {
	// 使用正则表达式查找单词边界
	camelRegexp := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	s = camelRegexp.ReplaceAllString(s, "${1}-${2}")

	// 将非字母数字字符替换为短横线
	spaceRegexp := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = spaceRegexp.ReplaceAllString(s, "-")

	return strings.ToLower(strings.Trim(s, "-"))
}

// PascalCase 将字符串转换为帕斯卡格式(首字母大写的驼峰格式)
func PascalCase(s string) string {
	camel := CamelCase(s)
	if len(camel) == 0 {
		return camel
	}
	return strings.ToUpper(camel[:1]) + camel[1:]
}

// SnakeCase 将字符串转换为蛇形格式
func SnakeCase(s string) string {
	// 处理驼峰格式
	camelRegexp := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	s = camelRegexp.ReplaceAllString(s, "${1}_${2}")

	// 将非字母数字字符替换为下划线
	spaceRegexp := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = spaceRegexp.ReplaceAllString(s, "_")

	return strings.ToLower(strings.Trim(s, "_"))
}

// Set operations

// Difference 返回两个切片的差集（在第一个切片但不在第二个切片中的元素）
func Difference[T comparable](slice1, slice2 []T) []T {
	result := []T{}
	lookup := make(map[T]bool)

	for _, v := range slice2 {
		lookup[v] = true
	}

	for _, v := range slice1 {
		if !lookup[v] {
			result = append(result, v)
		}
	}

	return result
}

// Intersection 返回两个切片的交集（同时在两个切片中的元素）
func Intersection[T comparable](slice1, slice2 []T) []T {
	result := []T{}
	lookup1 := make(map[T]bool)
	lookup2 := make(map[T]bool)

	for _, v := range slice1 {
		lookup1[v] = true
	}

	for _, v := range slice2 {
		lookup2[v] = true
		if lookup1[v] {
			result = append(result, v)
		}
	}

	return result
}

// UnionSlice 返回两个切片的并集（两个切片中所有的唯一元素）
func UnionSlice[T comparable](slice1, slice2 []T) []T {
	lookup := make(map[T]bool)
	result := []T{}

	for _, v := range slice1 {
		if !lookup[v] {
			lookup[v] = true
			result = append(result, v)
		}
	}

	for _, v := range slice2 {
		if !lookup[v] {
			lookup[v] = true
			result = append(result, v)
		}
	}

	return result
}

// SymmetricDifference 返回两个切片的对称差集（只在一个切片中的元素）
func SymmetricDifference[T comparable](slice1, slice2 []T) []T {
	union := UnionSlice(slice1, slice2)
	intersection := Intersection(slice1, slice2)

	return Difference(union, intersection)
}

// 数值处理函数

// Range 生成从start到end的数字序列，步长为step
func Range(start, end, step int) []int {
	if step == 0 {
		return []int{}
	}

	if (step > 0 && start >= end) || (step < 0 && start <= end) {
		return []int{}
	}

	var result []int
	for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
		result = append(result, i)
	}

	return result
}

// Sum 计算数字切片的总和
func Sum[N Number](slice []N) N {
	var sum N
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Avg 计算数字切片的平均值
func Avg[N Number](slice []N) float64 {
	if len(slice) == 0 {
		return 0
	}

	var sum N
	for _, v := range slice {
		sum += v
	}

	return float64(sum) / float64(len(slice))
}

// Min 返回数字切片中的最小值
func Min[N Number](slice []N) N {
	if len(slice) == 0 {
		var zero N
		return zero
	}

	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return min
}

// Max 返回数字切片中的最大值
func Max[N Number](slice []N) N {
	if len(slice) == 0 {
		var zero N
		return zero
	}

	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}

// Clamp 限制值在指定范围内
func Clamp[N Number](value, min, max N) N {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// LinearMap 将值从一个范围线性映射到另一个范围
func LinearMap[N Number](value, fromLow, fromHigh, toLow, toHigh N) N {
	if fromHigh == fromLow {
		return toLow
	}

	return toLow + (value-fromLow)*(toHigh-toLow)/(fromHigh-fromLow)
}

// Array/Slice utilities

// Compact 去除切片中的假值（零值）
func Compact[T comparable](slice []T) []T {
	var result []T
	var zero T

	for _, v := range slice {
		if v != zero {
			result = append(result, v)
		}
	}

	return result
}

// FlattenDeep 递归扁平化多维切片
func FlattenDeep(slice interface{}) []interface{} {
	result := []interface{}{}
	flattenRecursive(slice, &result)
	return result
}

func flattenRecursive(item interface{}, result *[]interface{}) {
	// 使用反射来处理嵌套切片，但这里简化实现，仅处理常见的多层嵌套
	switch v := item.(type) {
	case []interface{}:
		for _, subItem := range v {
			flattenRecursive(subItem, result)
		}
	case []int:
		for _, subItem := range v {
			*result = append(*result, subItem)
		}
	case []string:
		for _, subItem := range v {
			*result = append(*result, subItem)
		}
	case []float64:
		for _, subItem := range v {
			*result = append(*result, subItem)
		}
	default:
		*result = append(*result, v)
	}
}

// CompactBy 使用条件函数去除切片中的元素
func CompactBy[T any](slice []T, predicate func(T) bool) []T {
	var result []T

	for _, v := range slice {
		if !predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

// Uniq 返回去重后的切片
func Uniq[T comparable](slice []T) []T {
	lookup := make(map[T]bool)
	var result []T

	for _, v := range slice {
		if !lookup[v] {
			lookup[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Initial 返回除最后一个元素外的所有元素
func Initial[T any](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	return slice[:len(slice)-1]
}

// Tail 返回除第一个元素外的所有元素
func Tail[T any](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	return slice[1:]
}

// TakeRight 从右边开始取n个元素
func TakeRight[T any](slice []T, n int) []T {
	if n > len(slice) {
		n = len(slice)
	}
	if n < 0 {
		n = 0
	}
	return slice[len(slice)-n:]
}

// DropRight 从右边开始丢弃n个元素
func DropRight[T any](slice []T, n int) []T {
	if n > len(slice) {
		n = len(slice)
	}
	if n < 0 {
		n = 0
	}
	return slice[:len(slice)-n]
}

// SliceToMap 将切片转换为映射
func SliceToMap[T any, K comparable, V any](slice []T, keyFunc func(T) (K, V)) map[K]V {
	result := make(map[K]V)
	for _, item := range slice {
		k, v := keyFunc(item)
		result[k] = v
	}
	return result
}

// FindLast 查找满足条件的最后一个元素
func FindLast[T any](slice []T, predicate func(T) bool) (T, bool) {
	var zero T
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return slice[i], true
		}
	}
	return zero, false
}

// FindLastIndex 查找满足条件的最后一个元素的索引
func FindLastIndex[T any](slice []T, predicate func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return i
		}
	}
	return -1
}

// Object/Struct helpers (using map[string]interface{} as generic object)

// Get 获取嵌套对象中的值（类似JavaScript的get）
func Get(obj map[string]interface{}, path string, defaultValue interface{}) interface{} {
	keys := strings.Split(path, ".")
	current := obj

	for i, key := range keys {
		value, exists := current[key]
		if !exists {
			return defaultValue
		}

		if i == len(keys)-1 {
			return value
		}

		// 如果下一层是map[string]interface{}，继续查找
		if next, ok := value.(map[string]interface{}); ok {
			current = next
		} else {
			return defaultValue
		}
	}

	return defaultValue
}

// Pick 从对象中选择指定的键
func Pick(obj map[string]interface{}, keys []string) map[string]interface{} {
	result := make(map[string]interface{})
	keyMap := make(map[string]bool)

	for _, k := range keys {
		keyMap[k] = true
	}

	for k, v := range obj {
		if keyMap[k] {
			result[k] = v
		}
	}

	return result
}

// Omit 从对象中排除指定的键
func Omit(obj map[string]interface{}, keys []string) map[string]interface{} {
	result := make(map[string]interface{})
	keyMap := make(map[string]bool)

	for _, k := range keys {
		keyMap[k] = true
	}

	for k, v := range obj {
		if !keyMap[k] {
			result[k] = v
		}
	}

	return result
}

// Pluck 从对象数组中提取指定字段的值
func Pluck(slice []map[string]interface{}, field string) []interface{} {
	result := make([]interface{}, len(slice))

	for i, obj := range slice {
		result[i] = obj[field]
	}

	return result
}
