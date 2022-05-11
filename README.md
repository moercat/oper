# oper

The most suitable operation package for the job

最符合工作需要的操作包

![example](https://img.shields.io/badge/License-MIT-green)
![example](https://img.shields.io/badge/Go-1.16-blue)

## 目的

本项目采用了Golang 来制作 Python 中比较流行的组件，作为迅捷开发的工具，减少重复造轮子的时间

## 前言 

项目基于 https://github.com/spf13/cast 大佬所创建的类型转置模块，后期将会用泛型重构

[cast](https://github.com/spf13/cast): safe and easy casting from one type to another in Go ;在 Go 中安全且轻松地从一种类型转换为另一种类型

## 简便操作

将 python list 通过 go 实现,

1. In(fat, sub interface{}) bool
2. Index(fat, idx interface{}) int 不存在则为 -1
3. Extend(fat, sub interface{}) interface{}
4. Pop(fat interface{}, idx int) interface{}
5. Remove(fat interface{}, value interface{}) interface{}
6. Append(fat interface{}, value interface{}) interface{}
7. Insert(fat interface{}, idx int, value interface{}) interface{}
8. Count(fat interface{}, value interface{}) (count int)
9. Equal(fat, sub interface{}) bool

将 python set 通过 go 实现,


1. Set(fat interface{}) (sub interface{})
2. Add(fat interface{}, value interface{}) (sub interface{})
3. Union(fat interface{}, set ...interface{}) (sub interface{})
4. Update(fat interface{}, set interface{}) (sub interface{})

## 项目使用例子

```go
package main

func main() {

	//list := []int{
	//	1, 2, 3,
	//}
	// param Slice sub
	// return bool if exist true 
	// exist := oper.In(list, 1)
	// exist true list is  []int{ 1, 2, 3, } has 1
	// fmt.Println(exist)

}

```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod tidy // 自动管理
go run main.go 
```
