package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Message struct {
    data string
}

func subscribe(channel chan Message, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
			select {
			case channel <- Message{data: fmt.Sprintf("输入数字 %d", i)}:
			case <-ctx.Done(): // 使用 context 来处理超时
					fmt.Println("发送消息超时")
					return // 退出循环
			}
	}
}

func publisher(channel chan Message, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for message := range channel { // range 循环会自动处理 channel 关闭的情况
			fmt.Println("结果：", message)
	}
}

func exampleSub() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    var channel = make(chan Message, 2)
    wg := sync.WaitGroup{}
    wg.Add(2)

    go publisher(channel, ctx, &wg)
    go subscribe(channel, ctx, &wg)


    wg.Wait()
    fmt.Println("over")
}