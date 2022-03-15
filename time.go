package oper

import (
	"github.com/spf13/cast"
	"time"
)

func Parse(fat interface{}) string {

	fats := cast.ToTime(fat).In(time.Local)

	return fats.Format("2006-01-02 15:04:05")
}

// TodayDate time.Now().Format("2006-01-02")
func TodayDate() string {
	return time.Now().Format("2006-01-02")
}

// YesterdayDate time.Now().AddDate(0, 0, -1).Format("2006-01-02")
func YesterdayDate() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}
