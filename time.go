package oper

import (
	"github.com/spf13/cast"
	"time"
)

func Parse(fat interface{}) string {

	fats := cast.ToTime(fat).In(time.Local)

	return fats.Format("2006-01-02 15:04:05")
}

// Today time.Now().Format("2006-01-02")
func Today() string {
	return time.Now().Format("2006-01-02")
}

// Yesterday time.Now().AddDate(0, 0, -1).Format("2006-01-02")
func Yesterday() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}
