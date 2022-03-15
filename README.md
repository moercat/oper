# oper
opera package

# 简便操作
将 python list 通过 go 实现,
基于 https://github.com/spf13/cast 大佬所创建的类型转置模块

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
基于 https://github.com/spf13/cast 大佬所创建的类型转置模块

1. Set(fat interface{}) (sub interface{}) 
2. Add(fat interface{}, value interface{}) (sub interface{})
3. Union(fat interface{}, set ...interface{}) (sub interface{})
4. Update(fat interface{}, set interface{}) (sub interface{})
