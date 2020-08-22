####go struct能不能比较
`不同类型的不能比较，相同类型的可以比较`

#### go defer顺序
` 先入后出`

####GoRoot 和 GoPath 有什么用?
` goroot：go的安装路径；gopath：项目所在目录`

####go的new和make区别
`new：仅申请内存
 make：申请内存并初始化`
 
####go 如何实现长链接
`go 默认直接长链接，服务端设置超时时间即可`

####channel关闭阻塞问题
    
    channel 读写都是阻塞的，写设置缓冲区可以不在缓冲区范围内不阻塞
   
####CSP并发模型

    go的csp模型是基于channel实现的，两个goroutine通过channel来共享内存