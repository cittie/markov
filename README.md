# markov
- 自己弄着玩的马尔科夫链实现

## 概述
- 用int64代表权重实现的马尔科夫链
- 譬如下面这样的表

|  |  |  |     |
|---|---|---|-----|
|0|1|2| 4   |
|1|0|2| 4   |
|1|2|0| 4   |
|4|2|1| 0   |

- 表示下次事件如果从第一行起跳
  - 有1/7概率跳往第2行
  - 2/7概率跳往第3行
  - 4/7概率跳往第4行

## 代码示例
```go
// 初始化
ch := NewChain()
ch.AddStatus(NewStatus("line1", []int64{1}), []int64{})
ch.AddStatus(NewStatus("line2", []int64{1, 1}), []int64{2})
ch.AddStatus(NewStatus("line3", []int64{0, 3, 3}), []int64{4, 5})
ch.AddStatus(NewStatus("line4", []int64{4, 3, 2, 1}), []int64{9, 6, 3})

// 设置开始的序列
ch.SetCurStatusIdx(0)

// 随机跳跃
for i := 0; i < 20; i++ {
    n := ch.Next()
    fmt.Println(n)
}

fmt.Println(ch.CurStatusIdx())

```

