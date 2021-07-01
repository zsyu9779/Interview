## 同步原语和锁

Go语言在sync包中提供了用于同步的一些基本原语，包括cond，map，mutex，once，pool，rwmutex，waitgroup等，下文就针对

几种同步原语和包括在内的几种锁进行详细介绍

### Mutex

go语言中的Mutex结构体由两个字段state和sema构成，其中state表示当前互斥锁的状态，而sema是用于控制锁状态的信号量。

```go
type Mutex struct {
   state int32
   sema  uint32
}
```

