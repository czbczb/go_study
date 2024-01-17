package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readJsonFile() {
	var path = "test.json"
	 data, err := os.ReadFile(path)

	 if(err != nil) {
		fmt.Println("read json file err: ", err)
	 }

	 jsonData := make(map[string]interface{})
	 err = json.Unmarshal(data, &jsonData)

	 if(err != nil) {
		fmt.Println("unmarshal json file err: ", err)
	 }
	 for k, v := range jsonData {
		fmt.Println(k,v)
	 }
}
