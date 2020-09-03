### http和grpc的微服务框架demo

主要利用go监听http端口
利用consul用于服务的注册和发现

#### 可靠性
模拟了定时崩溃的程序，http服务的可靠性
当链接服务超时时，重新获取健康的服务


#### 缺少的组件

1.数据库/组件
2.redis库/组件
3.结构化日志
4.服务监控和上报