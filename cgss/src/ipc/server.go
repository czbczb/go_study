package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string `json:"method`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func(server *IpcServer) Connect() chan string {
	session := make(chan string)

	go func(c chan string) {
		for {
			request := <- c
			if request == "close" ||  request == "" {   // 关闭该链接
				break
			}
			var req Request

			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("InValid request format: ", request)
				c <- string("")
				continue
			}
			res := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(res)
			if err != nil {
				fmt.Println("handle request failed error: ", err)
				c <- string("")
				continue
			}

			c <- string(b)  // 返回结果  往session里面加一个相应结果
		}
	}(session)

	return session
}