# SkyWalking安装部署与集成

### 一、安装部署

#### 1.官网下载安装包

https://skywalking.apache.org/downloads/

![image-20230601161536483](/Users/zhangshiyu/Library/Application Support/typora-user-images/image-20230601161536483.png)

#### 2.下载tar包，解压

```shell
tar -zxvf apache-skywalking-apm-9.4.0.tar.gz
```

解压后文件目录为

![image-20230601162235270](/Users/zhangshiyu/Library/Application Support/typora-user-images/image-20230601162235270.png)



**Notice：！！！** skywalking运行依赖java的JDK11及以上版本 如果本机的java版本在11以下 先升级java版本到11以上

通过以下命令确定java版本

```shell
java -version
```

#### 3.对服务进行配置（可以不配 先run起来）

进入config目录 找到application.yml

```shell
vim application.yml
```

打开配置文件，找到如下段落

```yaml
...
storage:
  selector: ${SW_STORAGE:h2}
  elasticsearch:
    namespace: ${SW_NAMESPACE:""}
    clusterNodes: ${SW_STORAGE_ES_CLUSTER_NODES:localhost:9200}
    protocol: ${SW_STORAGE_ES_HTTP_PROTOCOL:"http"}
    connectTimeout: ${SW_STORAGE_ES_CONNECT_TIMEOUT:3000}
    socketTimeout: ${SW_STORAGE_ES_SOCKET_TIMEOUT:30000}
    responseTimeout: ${SW_STORAGE_ES_RESPONSE_TIMEOUT:15000}
    numHttpClientThread: ${SW_STORAGE_ES_NUM_HTTP_CLIENT_THREAD:0}
    user: ${SW_ES_USER:""}
    password: ${SW_ES_PASSWORD:""}
    trustStorePath: ${SW_STORAGE_ES_SSL_JKS_PATH:""}
    trustStorePass: ${SW_STORAGE_ES_SSL_JKS_PASS:""}
    secretsManagementFile: ${SW_ES_SECRETS_MANAGEMENT_FILE:""} # Secrets management file in the properties format includes the username, password, which are managed by 3rd party tool.
    ......
```

配置文件的形式是${AAAA:BBB}，其中 AAAA是需要配置的环境变量 BBB是默认配置

可以看到，skywalking默认的存储方式是基于java的轻量级数据库h2（所以不配也可以run起来）

下面关于elasticsearch的设置 即通过SW_NAMESPACE， SW_STORAGE_ES_CLUSTER_NODES，SW_STORAGE_ES_HTTP_PROTOCOL等环境变量配置

由于测试环境本机装有es 地址是localhost:9200，所以只需简单配置SW_STORAGE环境变量 使skywalking的存储方式变为ES即可

```shell
export SW_STORAGE=elasticsearch
```



配置文件中还有一些skywalking启动的默认配置

```yaml
 default:
    # Mixed: Receive agent data, Level 1 aggregate, Level 2 aggregate
    # Receiver: Receive agent data, Level 1 aggregate
    # Aggregator: Level 2 aggregate
    role: ${SW_CORE_ROLE:Mixed} # Mixed/Receiver/Aggregator
    restHost: ${SW_CORE_REST_HOST:0.0.0.0}
    restPort: ${SW_CORE_REST_PORT:12800}
    restContextPath: ${SW_CORE_REST_CONTEXT_PATH:/}
    restAcceptQueueSize: ${SW_CORE_REST_QUEUE_SIZE:0}
    httpMaxRequestHeaderSize: ${SW_CORE_HTTP_MAX_REQUEST_HEADER_SIZE:8192}
    gRPCHost: ${SW_CORE_GRPC_HOST:0.0.0.0}
    gRPCPort: ${SW_CORE_GRPC_PORT:11800}
    maxConcurrentCallsPerConnection: ${SW_CORE_GRPC_MAX_CONCURRENT_CALL:0}
    maxMessageSize: ${SW_CORE_GRPC_MAX_MESSAGE_SIZE:0}
    gRPCThreadPoolQueueSize: ${SW_CORE_GRPC_POOL_QUEUE_SIZE:-1}
    gRPCThreadPoolSize: ${SW_CORE_GRPC_THREAD_POOL_SIZE:-1}
    gRPCSslEnabled: ${SW_CORE_GRPC_SSL_ENABLED:false
    gRPCSslKeyPath: ${SW_CORE_GRPC_SSL_KEY_PATH:""}
    gRPCSslCertChainPath: ${SW_CORE_GRPC_SSL_CERT_CHAIN_PATH:""}
    gRPCSslTrustedCAPath: ${SW_CORE_GRPC_SSL_TRUSTED_CA_PATH:""}
```

如果使用grpc上报方式 关注 SW_CORE_GRPC_HOST和SW_CORE_GRPC_PORT 这是集成的时候的上报地址，如果需要http方式上报则需关注SW_CORE_REST_HOST和SW_CORE_REST_PORT

同理 在skywalking根目录下有个webapp目录 里面同样有个application.yml 和UI服务相关的配置在里面 容量不大 可以下来自行关注一下

#### 4.启动apm服务&UI界面

接下来进入 bin目录，执行启动命令

```shell
sh oapService.sh #启动apm服务
sh webappService.sh #启动UI界面
```

下一步进入根目录下log文件 关注apm服务端和webapp有无启动错误日志，如果正常 那么恭喜你安装成功

在浏览器打开http://localhost:8080/就可以看到skywalking的UI界面

![image-20230601170917386](/Users/zhangshiyu/Library/Application Support/typora-user-images/image-20230601170917386.png)

### 二、服务端接入（go-kit为例）

#### 1.名词解释

##### 1.span

各业务端在进行链路追踪的时候可以创建span 这样在UI界面上就会显示由一组或者多组span组成的span树 用来监控一条链路的行为

在skywalking中 业务端可以创建的span主要有三种：

- EntrySpan

```go
func (t *Tracer) CreateEntrySpan(ctx context.Context, operationName string, extractor propagation.Extractor) (s Span, nCtx context.Context, err error)
```

使用这个方法 创建的是 Entry Span，代表的是入站操作，也就是服务接收到的请求。比如说，一个 HTTP 请求到达你的服务，并开始处理这个请求，那么就应该在这时创建一个 Entry Span。

- LocalSpan



