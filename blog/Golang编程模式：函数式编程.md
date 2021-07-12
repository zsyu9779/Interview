## Golang编程模式(一) Functional Options

#### 应用场景举例：

在编程中，我们经常需要对一个对象（或是业务实体）进行相关的配置。比如下面这个业务实体

```go
type Server struct {
    Addr     string
    Port     int
    Protocol string
    Timeout  time.Duration
    MaxConns int
    TLS      *tls.Config
}
```

在这个 Server 对象中，我们可以看到：

要有侦听的 IP 地址 Addr 和端口号 Port ，这两个配置选项是必填的

还有协议 Protocol 、 Timeout 和MaxConns 字段，这几个字段是不能为空的，但是有默认值的，比如，协议是 TCP，超时30秒 和 最大链接数1024个

还有一个 TLS ，这个是安全链接，需要配置相关的证书和私钥。这个是可以为空的。

针对这样的配置，需要多种不同的函数签名创建不同配置的Server，例如：

```go
func NewDefaultServer(addr string, port int) (*Server, error) {
  return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
}

func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
  return &Server{addr, port, "tcp", 30 * time.Second, 100, tls}, nil
}

func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
  return &Server{addr, port, "tcp", timeout, 100, nil}, nil
}

func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls *tls.Config) (*Server, error) {
  return &Server{addr, port, "tcp", 30 * time.Second, maxconns, tls}, nil
}
```

就很蠢🤔

**有三种解决办法**

第一种比较简单但不推荐：引入一个Config{},将非必填的字段放到Config里，传入时按需赋值字段(也可以传入空struct，所以Server函数中要有相应的判空操作)

### Builder设计模式

（java开发中较常见）

第二种方法引入一个ServerBuilder{}

```go

//使用一个builder类来做包装
type ServerBuilder struct {
  Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
  sb.Server.Addr = addr
  sb.Server.Port = port
  //其它代码设置其它成员的默认值
  return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
  sb.Server.Protocol = protocol 
  return sb
}

func (sb *ServerBuilder) WithMaxConn( maxconn int) *ServerBuilder {
  sb.Server.MaxConns = maxconn
  return sb
}

func (sb *ServerBuilder) WithTimeOut( timeout time.Duration) *ServerBuilder {
  sb.Server.Timeout = timeout
  return sb
}

func (sb *ServerBuilder) WithTLS( tls *tls.Config) *ServerBuilder {
  sb.Server.TLS = tls
  return sb
}

func (sb *ServerBuilder) Build() (Server) {
  return  sb.Server
}
```

通过组合式编程(链式调用)，自由组合非必填参数：

```go
sb := ServerBuilder{}
server, err := sb.Create("127.0.0.1", 8080).
  WithProtocol("udp").
  WithMaxConn(1024).
  WithTimeOut(30*time.Second).
  Build()
```

### 函数式编程

golang style

首先定义一个函数类型变量

```go
type Option func(*Server)
```

然后，使用函数式编程定义以下一组函数

```go
func Protocol(p string) Option {
    return func(s *Server) {
        s.Protocol = p
    }
}
func Timeout(timeout time.Duration) Option {
    return func(s *Server) {
        s.Timeout = timeout
    }
}
func MaxConns(maxconns int) Option {
    return func(s *Server) {
        s.MaxConns = maxconns
    }
}
func TLS(tls *tls.Config) Option {
    return func(s *Server) {
        s.TLS = tls
    }
}
```

这组函数传入相应字段参数，返回一个用于设置*Server成员的匿名函数

例如调用MaxConns(30)时，返回值为匿名函数：

```go
func(s *Server){
  s.MaxConns = 30
}
```

接下来就可以定义一个初始化Server对象的函数NewServer(),入参中包含一个可变参数options，可以传入多个上面的函数参数，然后使用一个for-loop来设置Server对象。

```go

func NewServer(addr string, port int, options ...Option) (*Server, error) {

  srv := Server{
    Addr:     addr,
    Port:     port,
    Protocol: "tcp",
    Timeout:  30 * time.Second,
    MaxConns: 1000,
    TLS:      nil,
  }
  for _, option := range options {
    option(&srv)//执行函数参数，设置相应字段
  }
  //...
  return &srv, nil
}
```

于是，创建一个Server对象时，使用下面这样的代码：

```go
s1, _ := NewServer("localhost", 1024)
s2, _ := NewServer("localhost", 2048, Protocol("udp"))
s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
```

使用以上方式，不需引入Config{}或者使用Builder，代码优雅，提高了代码的可扩展性和可维护性，类似于可插拔式中间件的设计带来了高度的可配置化，使新使用者易于上手。

