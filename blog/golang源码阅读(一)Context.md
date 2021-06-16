## 上下文Context

上下文  Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体。上下文与 Goroutine 有比较密切的关系。

### context.Context接口

需要实现四个方法

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

##### Deadline 

- 返回 context.Context 被取消的时间，也就是完成工作的截止日期，如果未设置截止日期，则返回ok==false；

#####  Done

- 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，如果context永远无法取消（？）则可能返回nil，多次调用 Done 方法会返回同一个 Channel；channel的关闭可能会异步发生

##### Err 

- 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；

- 如果 context.Context 被取消，会返回 Canceled 错误；

- 如果 context.Context 超时，会返回 DeadlineExceeded 错误；

##### Value

- 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；

### 设计原理

### Context的相关操作

#### 默认上下文

context包中最常用的是context.Background context.TODO,两个方法会返回预先初始化好的私有变量

```go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}
```

emptyCtx永远不会取消，没有值，也没有截止日期。它不是struct {}，因为此类型的变量必须具有不同的地址。

emptyCtx通过空方法实现了context.Context接口的所有方法，他没有任何功能

Background和TODO都会返回一个emptyCtx，不同之处在于（其实只是语义上略有不同，实际上二者互为别名）

- Background通常由main方法，初始化和单元测试使用，并且用作传入请求(request)的顶级上下文
- TODO主要用于当不清楚要使用哪个Context或尚不可用时（因为尚未扩展周围的函数以接受Context参数），

#### 取消信号

[`context.WithCancel`](https://draveness.me/golang/tree/context.WithCancel) 函数能够从 [`context.Context`](https://draveness.me/golang/tree/context.Context) 中衍生出一个新的子上下文并返回用于取消该上下文的函数。一旦我们执行返回的取消函数，当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号。

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}

func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}
```

- WithCancel方法通过newCancelCtx返回一个cancelCtx结构体

- propagateCancel方法：会构建父子上下文之间的关联，当父上下文被取消时，子上下文也会被取消：

  ```go
  func propagateCancel(parent Context, child canceler) {
     done := parent.Done()
     if done == nil {
        return // parent is never canceled
     }
  
     select {
     case <-done:
        // parent is already canceled
        child.cancel(false, parent.Err())
        return
     default:
     }
  
     if p, ok := parentCancelCtx(parent); ok {
        p.mu.Lock()
        if p.err != nil {
           // parent has already been canceled
           child.cancel(false, p.err)
        } else {
           if p.children == nil {
              p.children = make(map[canceler]struct{})
           }
           p.children[child] = struct{}{}
        }
        p.mu.Unlock()
     } else {
        atomic.AddInt32(&goroutines, +1)
        go func() {
           select {
           case <-parent.Done():
              child.cancel(false, parent.Err())
           case <-child.Done():
           }
        }()
     }
  }
  ```

上述函数总共与父上下文相关的三种不同的情况：

1. 当 `parent.Done() == nil`，也就是 `parent` 不会触发取消事件时，当前函数会直接返回；
2. 当child的继承链包含可以取消的上下文时，会判断parent是否已经触发了取消信号；
   - 如果已经被取消，`child` 会立刻被取消；
   - 如果没有被取消，`child` 会被加入 `parent` 的 `children` 列表中，等待 `parent` 释放取消信号；
3. 当父上下文是开发者自定义的类型、实现了`context.Context`接口并在Done()方法中返回了非空的管道时；
   1. 运行一个新的 Goroutine 同时监听 `parent.Done()` 和 `child.Done()` 两个 Channel；
   2. 在 `parent.Done()` 关闭时调用 `child.cancel` 取消子上下文；



cancelCtx：

```go
type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     chan struct{}         // created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
}
```

canceler是可以直接取消的Context类型，canceler接口：

```go
type canceler interface {
   cancel(removeFromParent bool, err error)
   Done() <-chan struct{}
}
```

cancelCtx本身也继承了canceler接口和context.Context接口：

```go
func (c *cancelCtx) Value(key interface{}) interface{} {
   if key == &cancelCtxKey {
      return c
   }
   return c.Context.Value(key)
}

func (c *cancelCtx) Done() <-chan struct{} {
   c.mu.Lock()
   if c.done == nil {
      c.done = make(chan struct{})
   }
   d := c.done
   c.mu.Unlock()
   return d
}

func (c *cancelCtx) Err() error {
   c.mu.Lock()
   err := c.err
   c.mu.Unlock()
   return err
}

func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}
	c.err = err
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}
	for child := range c.children {
		// NOTE: acquiring the child's lock while holding parent's lock.
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		removeChild(c.Context, c)
	}
}
```

在cancel方法里 

1. 关闭cancelCtx的done channel 

2. 给cancelCtx的err赋值

3. 遍历cancelCtx的children(子Ctx)并链式调用其cancel方法

4. 若removeFromParent为true则代表当前Ctx存在父级Ctx，则从父级Ctx的children中删除当前Ctx：

   ```
   func removeChild(parent Context, child canceler) {
      p, ok := parentCancelCtx(parent)
      if !ok {
         return
      }
      p.mu.Lock()
      if p.children != nil {
         delete(p.children, child)
      }
      p.mu.Unlock()
   }
   ```

除了 context.WithCancel 之外，context 包中的另外两个函数 context.WithDeadline 和 context.WithTimeout 也都能创建可以被取消的计时器上下文 context.timerCtx

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
   if parent == nil {
      panic("cannot create context from nil parent")
   }
   if cur, ok := parent.Deadline(); ok && cur.Before(d) {
      // The current deadline is already sooner than the new one.
      return WithCancel(parent)
   }
   c := &timerCtx{
      cancelCtx: newCancelCtx(parent),
      deadline:  d,
   }
   propagateCancel(parent, c)
   dur := time.Until(d)
   if dur <= 0 {
      c.cancel(true, DeadlineExceeded) // deadline has already passed
      return c, func() { c.cancel(false, Canceled) }
   }
   c.mu.Lock()
   defer c.mu.Unlock()
   if c.err == nil {
      c.timer = time.AfterFunc(dur, func() {
         c.cancel(true, DeadlineExceeded)
      })
   }
   return c, func() { c.cancel(true, Canceled) }
}
```

WithDeadline方法通过比较父context的截止时间和当前context的截止时间，算出context存活时间，通过调用time.AfterFunc匿名函数定时取消context

#### 传值方法

父context调用WithValue方法可以创建一个子上下文用来传值，子context的类型为valueCtx：

```go
type valueCtx struct {
   Context
   key, val interface{}
}
```

valueCtx继承了Context，另携带了一个键值对，重写了Value方法：

```go
func (c *valueCtx) Value(key interface{}) interface{} {
   if c.key == key {
      return c.val
   }
   return c.Context.Value(key)
}
```

如果传入的键与context中存储的不相符则从父上下文中查找该键对应的值，直到某个父context返回nil或者查找到对应的值