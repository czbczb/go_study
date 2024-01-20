package distinctReq

import (
	"fmt"
	"testing"
	"time"
)


func Test_distinctReq(t *testing.T) {
	bucket := NewTokenBucket(20, time.Second)
	for i := 0; i< 60; i++ {
		if bucket.Take() {
			fmt.Println("请求处理了")
		}else {
			fmt.Println("请求拒绝了")
		}
		time.Sleep(200 * time.Millisecond)
	}
}