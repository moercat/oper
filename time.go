package oper

import (
	"github.com/spf13/cast"
	"time"
)

// Parse 将输入转换为本地时间并格式化为 "YYYY-MM-DD HH:MM:SS" 格式
func Parse(fat interface{}) string {
	fats := cast.ToTime(fat).In(time.Local)
	return fats.Format("2006-01-02 15:04:05")
}

// Today 返回今天日期，格式为 "YYYY-MM-DD"
func Today() string {
	return time.Now().Format("2006-01-02")
}

// Yesterday 返回昨天日期，格式为 "YYYY-MM-DD"
func Yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}
