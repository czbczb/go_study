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

func UseSlice () {
	MySlice := make([]int, 5, 10)
	// MySlice[6] = 2
	fmt.Println("Slice:", MySlice)

	fmt.Println(cap(MySlice), len(MySlice))

}

func UseMap() {
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
}

func main() {

	// testSlice()

	// lock()

	// syncCond()   // error

	// selectCase()
	// useWaitGroup()
	// useAsyncRoutine()

	// useSyncRoutine()

	// useContext()

	UseSlice()
}










