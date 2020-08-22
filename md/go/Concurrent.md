###高并发、goroutine
####go 为什么能做到高并发

    goroutine是Go并行设计的核心
        - goroutine非常轻量级
        - 线程内自主控制，无需cpu上下文切换
        - 拥有自己的寄存器上下文和栈
        - PGM的调度模型也决定了go有更高的并发性能
        
####GPM调度模型
    
    1.操作系统眼里只有进程和线程
    2.goroutine 运行在线程上，分别有，GPM分别代表goroutine,processor和 machine
    3.G是goroutine,含有执行的上下文寄存器和栈信息
    4.P是processor,维护了G的列表和缓存、状态等信息
    5.M是 machine，M对接一个P，丢给线程去处理
    6.为了保持P队列消耗不均匀，存在一个全局的队列。P队列执行G完会从其他P的一半，拿完从全局拿。
    7.协程是抢占式的，一些sleep、网络等io请求会阻塞goroutine的时候，会有协调则将请求阻塞到goroutine中，调度器就可以执行其他G了
    8.阻塞的G恢复后，会进入他的local P，不存在P则进入全局队列
    9.同一个P里面的G先进先出
    10.goroutine是抢占式调度，10ms后标识为可抢占的
    
####不使用锁来保证Golang的并发安全

    sync.mutex
    atomic.value
    redis/数据库等分布式锁
    chan 信号量传递消息控制并发，类似生产者消费者模型
