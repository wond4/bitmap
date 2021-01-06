# bitmap
使用 golang 实现的位图

## 说明
golang 中可以直接操作的最小内存单位为 byte（8 bit），如 bool、int8、uint8，这对于绝大多数场景是足够的。
但是当使用布隆过滤器或用筛法找素数时，这个精度会浪费大量内存，所以我们需要 bitmap 把最小内存单位延伸到 bit，
这可以节省 87.5% 的内存。

## 使用
```go
bm := NewBitmap(10) 
bm.Add(15)
bm.Add(0)
fmt.Println(bm.Has(0))
fmt.Println(bm.Has(1))
fmt.Println(bm)
bm.Del(0)
fmt.Println(bm.Has(0))
fmt.Println(bm)
```

## API 说明
### NewBitmap(length int) *Bitmap
得到一个新的 bitmap，length 为最大长度
> PS: 序号从 0 开始，如 NewBitmap(10) 可以放入 0-9 之间的数字

### (b *Bitmap) Add(num int) bool
往 bitmap 中添加一个数字，返回值表示是否添加成功（长度足够一定会成功）

### (b *Bitmap) Del(num int) bool bool
从 bitmap 中删除一个数字，返回值表示是否删除成功（长度足够一定会成功）

### (b *Bitmap) Len() int
返回 bitmap 最大长度

### (b *Bitmap) LenRaw() int
返回 bitmap 底层数组的长度（bytes）

### (b *Bitmap) String() string
返回 bitmap 二进制表示（注意过大的 bitmap 会产生性能问题）
