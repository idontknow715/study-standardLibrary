# 1. Go的并发模型是基于goroutine+channel的，但是sync提供了底层的并发原语，让我们能更细粒度的控制并发

# 2. 核心方法
## 2.1 sync.Mutex(互斥锁)
作用：
* 用来保证某段代码同一时刻只有一个goroutine能执行。
* 避免多个goroutine同时修改共享变量，导致数据错乱
方法：
* Lock(): 加锁（如果别的goroutine已经持有所，就会阻塞）
* Unlock(): 解锁（一定要成对使用，最好用defer）

## 2.2 sync.RwMutex(读写锁)
作用：
* 多个读可以同时进行，但写操作必须独占
* 适合读多写少的场景（比如缓存）

方法：
* RLock()/RUnLock(): 读锁
* Lock()/UnLock(): 写锁

## 2.3 sync.WaitGroup(等待组)
作用：
* 用来等待一组goroutine执行完成
方法：
* Add(n int): 添加计数
* Done(): 减少计数(常和defer搭配)
* Wait(): 阻塞，直到计数为0

## 2.4 sync.Once
作用：
保证某个操作只执行一次，常用与单例模式或只初始化一次资源

方法：
* Do(func()): 只执行一次函数

## 2.5 sync.Cond(条件变量)
作用：
* 用于等待/通知 场景
* 常和Mutex搭配，用来协调goroutine

方法：
* Wait(): 等待通知（会自动释放锁，并在收到信号时重新加锁）
* Signal(): 通知一个等待的goroutine
* BroadCast(): 通知所有等待的goroutine

## 3. 小练习
* 用sync.Mutex写一个并发安全计数器(100个goroutine同时+1)
* 用sync.RWMutex写一个并发安全字典(支持读写操作)
* 用sync.WaitGroup启动10个goroutine，每个打印自己的ID，最后打印done
* 用sync.Once实现一个单例初始化，比如只初始化数据库连接一次
* 用sync.Cond模拟一个生产者-消费者模型(生产者放数据，消费者等待数据)