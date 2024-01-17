package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func marshal() {
	// 创建一个 User 结构体
	user := User{
		Name: "John Doe",
		Age:  30,
	}

	// 序列化 User 结构体
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印序列化后的 JSON 数据
	fmt.Println(string(data))

	// 反序列化 JSON 数据
	var newUser User
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印反序列化后的 User 结构体
	fmt.Println(newUser)
}

func mutexLock() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	mutex.Lock()

	for i := 1; i <= 3; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)

			mutex.Lock()
			fmt.Println("Lock:", i)

			time.Sleep(time.Second)

			fmt.Println("Unlock:", i)
			mutex.Unlock()

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()

	wait.Wait()
}

func syncCond() {
	// 条件变量

}

func selectCase() error {
	ch := make(chan int)

	select {
	case ch <- 10: // 将值 10 发送到通道 ch
		fmt.Println("input: ", ch)
		return nil
	case value := <-ch: // 从通道 ch 接收值并赋给变量 value
		fmt.Println("value:", value)
		return nil
	default: // 在所有 `case` 都无法执行时执行
		fmt.Println("No value received")
		return nil
	}
}

func testSlice() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 删除 slice 中的第一个元素
	slice = append(slice[:0], slice[1:]...)

	// 打印 slice 中的所有元素
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func useWaitGroup() {
	// 创建一个 waitGroup
	var wg sync.WaitGroup

	// 启动 3 个协程
	for i := 0; i < 3; i++ {
		go func() {
			// 执行任务
			fmt.Println("协程", i, "开始执行")
			time.Sleep(time.Second)
			fmt.Println("协程", i, "执行结束")
			wg.Done()
		}()

		wg.Add(1)
	}

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有协程都已完成")
}

// 同步
func useSyncRoutine() {
	// 创建一个 channel
	ch := make(chan int, 1)
	// var wg sync.WaitGroup

	// 创建一个 goroutine 来发送数据
	go func() {
		// 发送数据
		ch <- 10
	}()
	// 创建一个 goroutine 来接收数据
	go func() {
		// 使用 select 语句检查 channel 是否有数据可用
		select {
		case x := <-ch:
			fmt.Println(x)
			break
		default:
			fmt.Println("No data available")
		}
	}()

	// 等待所有 goroutine 完成
	time.Sleep(time.Second)

}

// 异步
func useAsyncRoutine() {
	// 创建一个 channel

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ch := make(chan int)

	// 创建一个 goroutine 来发送数据
	go func() {
		// 发送数据
		ch <- 10
	}()

	// 创建一个 goroutine 来接收数据
	go func() {
	forLoop:
		for {
			// 使用 select 语句检查 channel 是否有数据可用
			select {
			case x := <-ch:
				fmt.Println(x)
				close(c)
				break forLoop
			default:
				fmt.Println("No data available")
			}
		}
	}()

	<-c
	// 主 goroutine 继续执行
	fmt.Println("Hello, world!")
}

func useContext() {
	// 创建一个带取消信号的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 创建一个 goroutine
	go func(ctx context.Context) {
		// 等待取消信号
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine 被取消")
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	// 等待 5 秒后取消 context
	time.Sleep(5 * time.Second)
	cancel()

	// 等待 goroutine 结束
	<-ctx.Done()
}

func useSyncPol() {
	// 创建一个协程池
	goroutinePool := &sync.Pool{
		// 存放空闲的协程
		// Capacity: 10,

		// 创建协程的函数
		// Factory: func() interface{} {
		// 	return make(chan func())
		// },
	}

	// 创建 10 个任务
	for i := 0; i < 10; i++ {
		// 从协程池中获取一个空闲的协程
		goroutine := goroutinePool.Get().(chan func())

		// 执行任务
		goroutine <- func() {
			fmt.Println("协程正在执行任务")
			time.Sleep(1 * time.Second)
		}
	}

	// 等待所有任务完成
	time.Sleep(10 * time.Second)
}

// 闭包
func useIncrement() {
	// 定义一个闭包
	increment := func(x int) int {
		return x + 1
	}

	// 调用闭包
	fmt.Println(increment(10))
}

// 装饰器
func useWrapper() {
	// 定义一个函数
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}

	// 定义一个装饰器
	logger := func(f func(name string)) func(name string) {
		return func(name string) {
			// 执行原始函数
			f(name)

			// 添加额外的功能
			fmt.Println("Log:", name)
		}
	}

	// 使用装饰器装饰函数
	greet = logger(greet)

	// 调用函数
	greet("John Doe")
}

type Singleton struct {
	Name string
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

func main() {

	// testSlice()

	// lock()

	// syncCond()   // error

	// selectCase()
	// useWaitGroup()
	// useAsyncRoutine()

	// useSyncRoutine()

	var configMap = map[string]string {
		"name": "configmap",
		"aws": "aws",
		"aliyun": "aliyun",
	}

	// for key, value := range configMap {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	configJson, err := json.Marshal(configMap)

	if(err != nil) {
		fmt.Print("marshal configjson error: ", err)
	}

	var configMap1 = make(map[string]string)
	err = json.Unmarshal(configJson, &configMap1)

	if(err != nil) {
		fmt.Print("unmarshal configjson error: ", err)
	}
	fmt.Print("configJson: ", configMap1)
	// useContext()
}


// go test -run ^Test_readFile

// 1、Golang中的大端序和小端序是什么？
// 答：大端序指数据的高位字节存储在内存的低地址处，小端序则相反，数据的低位字节存储在内存的低地址处。Golang默认采用小端序

// 2、Golang中的interface{}类型是什么？
// 答：interface{}类型是Golang中的空接口类型，可以表示任意类型的对象。可以使用type Assertion技术将interface{}转换为其他类型的对象。

// 3、内置函数new()
// 返回一个指针，指向分配的内存，该内存的值为T类型的零值

// 4、Golang中的make()和new()函数有什么区别？
// 答：make()函数用于创建slice、map、channel等引用类型的数据结构，返回的是该类型的引用；new()函数用于创建值类型的数据结构，并分配内存空间，返回的是指向该类型的指针。

// 5、Golang中的管道是什么？如何使用？
// 答：管道是Golang中一种由操作系统内核提供的机制，用于实现进程间通信。可以通过os.Pipe()函数创建一个管道，然后使用文件操作的方式进行读写。

// 6、Golang中的main()函数是否可以有返回值？
// 答：可以，main()函数可以有一个整型的返回值，表示程序的退出状态码。0表示程序正常退出，非0表示程序异常退出。

// 7、Golang中的反射机制是什么？如何使用？
// 答：反射机制是Golang中的一种特性，可以在运行时动态地获取变量的类型和值，并进行操作。可以使用reflect包中提供的相关方法和类型来实现反射操作

// Golang中的sync.Once是什么？

// 8、Golang中的协程和线程有什么区别？
// 答：Golang的协程是一种轻量级线程，相比于传统线程更加高效。在Golang中，每个协程只需要几KB的内存，并且由Go运行时自动管理调度，不需要手动创建和销毁线程。

// 9、Golang中的map是如何实现的？
// 答：Golang中的map采用哈希表来实现，具体实现方式为桶+链式法。即将所有键值对分配到一定数量的桶中，每个桶中都维护一个键值对链表。当需要访问某个键值对时，先通过哈希函数计算出其在哪个桶中，然后在该桶的链表中查找。

// 标准库sort包提供各种排序算法：快速排序、归并排序

// 10、Golang中的panic和recover语句是干什么用的？
// 答：panic和recover语句用于处理程序运行时的异常情况。当程序出现不可恢复的错误时，可以使用panic语句抛出异常。在上层函数中使用defer和recover语句可以捕获该异常，并进行相应的处理，从而避免程序崩溃

// 11、Golang中的字符串是如何实现的？
// 答：Golang中的字符串实际上是一个只读的字节数组片段，每个元素代表一个Unicode字符。由于字符串是只读的，所以在需要修改字符串内容时需要先将其转换成可修改的字节数组。

// 12、Golang中的init函数是什么？
// 答：init函数是一个特殊的函数，在程序运行之前会被自动执行。可以在该函数中进行全局变量初始化、注册驱动等操作。

// 13、Golang中的GC是什么？
// 答：GC（Garbage Collection）是Golang的垃圾回收机制，用于自动管理内存。在程序运行时，GC会定期扫描所有已分配的内存，将未被引用的内存自动回收。

// 重载实现：可变参数的函数

// 14、Golang中的字符串拼接有什么要注意的地方？
// 答：Golang中的字符串拼接可以使用"+"或者fmt.Sprintf函数。但是由于字符串是只读的，每次拼接都会创建一个新的字符串对象，并将原有字符串复制到新的对象中。因此，频繁的字符串拼接会造成性能问题。

// 15、Golang中如何进行内存池管理？
// 答：Golang标准库中提供了sync.Pool包，可以用于实现内存池管理。通过在goroutine中使用Pool.Get和Pool.Put函数，可以更加高效地分配和回收内存。




// 1、goroutine 利用多个cpu的条件            (1.17后不用再手动设置了)
// 1、设置环境变量 GOMAXPROCS = 16
// 2、执行runtime.GOMAXPROCS(16)
// 3、获取cpu的核数  runtime.NumCPU()

// 主动让出时间片：runtime.Gosched()




// 2、sync
// sync.mutex    sync.RWMutex(单写多读模式)
//    m.lock()      // 写    释放：m.Unlock()
// 	 m.RLock()     // 读    释放：m.RUnlock()

// var once sync.Once  
// once.Do(some func)  全局初始化某个方法，全局唯一










