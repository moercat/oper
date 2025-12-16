# oper

The most suitable operation package for the job (Supports Go 1.18+ with backward compatibility for features in newer Go versions)

最符合工作需要的操作包（支持Go 1.18+，向后兼容新版本Go中的功能）

![License](https://img.shields.io/badge/License-MIT-green)
![Go Version](https://img.shields.io/badge/Go-1.18-blue)

## 目的

本项目采用Golang实现Python中比较流行的列表、集合和时间等操作功能，作为快速开发的工具包，减少重复造轮子的时间。

## 特性

- 支持泛型的切片操作（如 In, Index, Extend, Pop, Remove, Append, Insert, Count, Equal）
- Python风格操作（如 Map映射, Filter过滤, Any检查, All检查）
- JavaScript风格数组方法（如 Find查找, FindIndex查找索引, Some部分满足, Every全部满足）
- Lodash风格实用程序（如 Chunk分块, Flatten扁平化, UniqBy条件去重, GroupBy分组）
- 额外字符串处理工具（如 Capitalize首字母大写, CamelCase驼峰转换, SnakeCase蛇形转换）- Go标准库中没有的便捷功能
- 额外集合操作（如 Set去重, Add添加元素, Union联合, Update更新, Difference差集, Intersection交集）- Go标准库中没有的切片集合操作
- 数值运算（如 Opera乘除法, Between范围判断, Range范围生成, Sum求和, Avg平均值, Min最小值, Max最大值）
- 额外数组实用工具（如 Compact压缩, Uniq去重, Initial非尾元素, Tail非首元素, TakeRight右边取值, DropRight右边丢弃）- Go标准库中没有的高级切片操作
- 时间处理（如 Parse解析时间, Today获取今天日期, Yesterday获取昨天日期）
- 对象助手函数（如 Get获取嵌套值, Pick选择键, Omit排除键）- Go标准库中没有的结构体/映射操作
- 通过 [spf13/cast](https://github.com/spf13/cast) 实现安全易用的类型转换

## 安装

```bash
go get github.com/moercat/oper
```

## 使用示例

### 切片操作

```go
package main

import (
    "fmt"
    "oper"
)

func main() {
    list := []int{1, 2, 3, 4, 2}
    
    // 检查元素是否存在
    exists := oper.In(list, 2)  // true
    fmt.Println("存在:", exists)
    
    // 查找元素索引
    index := oper.Index(list, 4)  // 3
    fmt.Println("索引:", index)
    
    // 计数
    count := oper.Count(list, 2)  // 2
    fmt.Println("计数:", count)
    
    // 删除元素
    newList := oper.Remove(list, 2)  // [1, 3, 4, 2] - 只删除第一个匹配项
    fmt.Println("删除后:", newList)
    
    // 添加元素
    extended := oper.Extend(list, []int{5, 6})  // [1, 2, 3, 4, 2, 5, 6]
    fmt.Println("扩展后:", extended)
    
    // 插入元素
    inserted := oper.Insert(list, 1, 99)  // [1, 99, 2, 3, 4, 2]
    fmt.Println("插入后:", inserted)

    // Python风格映射
    doubled := oper.Map(list, func(x int) int { return x * 2 })  // [2, 4, 6, 8, 4]
    fmt.Println("映射后:", doubled)

    // Python风格过滤
    evens := oper.Filter(list, func(x int) bool { return x%2 == 0 })  // [2, 4, 2]
    fmt.Println("偶数:", evens)

    // 检查是否有符合条件的元素
    hasEven := oper.Any(list, func(x int) bool { return x%2 == 0 })  // true
    fmt.Println("有偶数:", hasEven)

    // 检查是否所有元素都符合条件
    allEven := oper.All(list, func(x int) bool { return x%2 == 0 })  // false
    fmt.Println("全是偶数:", allEven)

    // JavaScript风格查找
    foundVal, foundIdx := oper.Find(list, func(x int) bool { return x > 3 })  // 4, 3
    fmt.Println("找到值和索引:", foundVal, foundIdx)

    // 查找符合条件的元素的索引
    foundIndex := oper.FindIndex(list, func(x int) bool { return x == 4 })  // 3
    fmt.Println("找到索引:", foundIndex)

    // 检查是否有元素满足条件
    hasSome := oper.Some(list, func(x int) bool { return x > 10 })  // false
    fmt.Println("有些元素大于10:", hasSome)

    // 检查是否所有元素都满足条件
    everyGTZero := oper.Every(list, func(x int) bool { return x > 0 })  // true
    fmt.Println("所有元素都大于0:", everyGTZero)

    // 分块
    chunks := oper.Chunk(list, 2)  // [[1, 2], [3, 4], [2]]
    fmt.Println("分块:", chunks)

    // 反转
    reversed := oper.Reverse(list)  // [2, 4, 3, 2, 1]
    fmt.Println("反转后:", reversed)

    // 取前n个元素
    firstThree := oper.Take(list, 3)  // [1, 2, 3]
    fmt.Println("前三个:", firstThree)

    // 丢弃前n个元素
    skipTwo := oper.Drop(list, 2)  // [3, 4, 2]
    fmt.Println("跳过前两个:", skipTwo)
}
```

### Python/JavaScript/Lodash风格操作

```go
package main

import (
    "fmt"
    "oper"
)

func main() {
    list := []int{1, 2, 3, 4, 5}

    // Map - 对每个元素执行变换函数
    squares := oper.Map(list, func(x int) int { return x * x })
    fmt.Println("平方:", squares)  // [1, 4, 9, 16, 25]

    // Filter - 过滤出满足条件的元素
    evens := oper.Filter(list, func(x int) bool { return x%2 == 0 })
    fmt.Println("偶数:", evens)  // [2, 4]

    // Any/Some - 检查是否有任意元素满足条件
    hasEven := oper.Any(list, func(x int) bool { return x%2 == 0 })
    fmt.Println("是否有偶数:", hasEven)  // true

    // All/Every - 检查是否所有元素都满足条件
    allPositive := oper.All(list, func(x int) bool { return x > 0 })
    fmt.Println("是否都为正数:", allPositive)  // true

    // Find/FindIndex - 查找满足条件的元素及其索引
    val, idx := oper.Find(list, func(x int) bool { return x > 3 })
    fmt.Printf("找到值: %d, 索引: %d\n", val, idx)  // 找到值: 4, 索引: 3

    // Chunk - 将切片分块
    chunkedList := oper.Chunk([]int{1, 2, 3, 4, 5, 6}, 2)
    fmt.Println("分块结果:", chunkedList)  // [[1 2] [3 4] [5 6]]

    // Flatten - 扁平化嵌套切片
    flattened := oper.Flatten([][]int{{1, 2}, {3, 4}, {5}})
    fmt.Println("扁平化结果:", flattened)  // [1 2 3 4 5]

    // GroupBy - 按条件分组
    groupResult := oper.GroupBy([]string{"apple", "ant", "banana", "berry"},
        func(s string) rune { return rune(s[0]) })
    fmt.Println("按首字母分组:", groupResult)

    // Contains - 检查是否包含某个元素
    containsThree := oper.Contains(list, 3)
    fmt.Println("包含3:", containsThree)  // true
}
```

### 集合操作

```go
package main

import (
    "fmt"
    "oper"
)

func main() {
    // 创建一个唯一元素的集合
    unique := oper.Set([]string{"a", "b", "c", "a"})  // ["a", "b", "c"]
    fmt.Println("去重后:", unique)

    // 添加元素到集合
    added := oper.Add([]string{"a", "b"}, "c")  // ["a", "b", "c"]
    fmt.Println("添加后:", added)

    // 联合多个集合
    union := oper.Union([]string{"a", "b"}, []string{"b", "c"}, []string{"c", "d"})  // ["a", "b", "c", "d"]
    fmt.Println("联合后:", union)

    // Lodash风格的去重 - 根据自定义函数去重
    people := []string{"alice", "bob", "charlie", "alex"}
    uniqueByFirstChar := oper.UniqBy(people, func(name string) string {
        return string(name[0]) // 按首字符去重
    })
    fmt.Println("按首字符去重:", uniqueByFirstChar)  // [alice bob charlie]
}
```

### 数值运算

```go
package main

import (
    "fmt"
    "oper"
)

func main() {
    // 数值乘法 (正数)
    result1 := oper.Opera("10", 5)  // "50"
    fmt.Println("10 * 5 =", result1)
    
    // 数值除法 (负数)
    result2 := oper.Opera("10", -2)  // "5"
    fmt.Println("10 / 2 =", result2)
    
    // 判断数值是否在范围内
    inRange, outRange, err := oper.Between(5, 1, 10)  // false, false, nil
    if err != nil {
        fmt.Println("错误:", err)
    } else {
        fmt.Printf("在范围内: %t, 不在范围内: %t\n", !outRange, inRange)
    }
}
```

### 时间操作

```go
package main

import (
    "fmt"
    "oper"
)

func main() {
    // 格式化时间
    formatted := oper.Parse("2023-01-01 12:30:45")  // "2023-01-01 12:30:45"
    fmt.Println("格式化时间:", formatted)
    
    // 获取今天日期
    today := oper.Today()  // 如 "2023-06-15"
    fmt.Println("今天:", today)
    
    // 获取昨天日期
    yesterday := oper.Yesterday()  // 如 "2023-06-14"
    fmt.Println("昨天:", yesterday)
}
```

## 泛型支持

本库利用Go 1.18+的泛型特性，提供了类型安全的操作：

```go
type Number interface {
    ~int8 | ~int16 | ~int | ~int32 | ~int64 |
    ~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64
}

type List interface {
    Number | ~string
}
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod tidy  # 自动管理依赖
go run main.go
```

## 测试

本项目包含完整的测试套件，覆盖所有功能：

```bash
go test -v
```

## 关于Go 1.21+的slices包

本库在Go 1.21+引入`slices`包之前就已经提供了许多类似的切片操作功能（如Contains, Index, Insert, Delete等）。
本库继续提供这些功能是为了保持对Go 1.18+版本的向后兼容性。
如果您使用Go 1.21+，也可以选择使用标准库的`slices`包。

## License

MIT