package ipc

import (
	"encoding/json"
	"fmt"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	client := server.Connect()
	return &IpcClient{client}
}

func (c *IpcClient) Call(method, params string) (resp Response, err error) {
	req  := &Request{method, params}

	var b []byte
	b, err = json.Marshal(req)

	if err != nil {
		fmt.Println("Error marshalling request:", err, method, params)
		return
	}
	str2 := string(b)
	fmt.Println(str2)
	c.conn <- string(b)
	str := <-c.conn   // 等待返回值
	var resp1 Response

	err = json.Unmarshal([]byte(str), &resp1)
	if err != nil {
		fmt.Println("Error unMarshall response", err, method, params)
	}
	return resp1, nil
}


func (c *IpcClient) Close() {
	// c.conn <- "close"
	close(c.conn)
}
