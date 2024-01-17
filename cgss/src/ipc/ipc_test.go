package ipc
import (
	"testing"
)

type EchoServer struct {

}

func (e EchoServer) Handle(method, params string) *Response {
	return &Response{ 
		Code: "ok",
		Body: "ECHO: " + method + " ~ " + params,
	}
}

func (e EchoServer) Name() string {
	return "EchoServer"
}

func Test_ipc (t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("getName", "From client 1")
	resp2, _ := client2.Call("getName", "From client 2")

	if resp1.Body != "ECHO: getName ~ From client 1" || resp2.Body != "ECHO: getName ~ From client 2"{
		t.Error("IpcClient.Call failed. resp1: ", resp1, "resp2: ", resp2.Body)
	}

	client1.Close()
	client2.Close()
}

