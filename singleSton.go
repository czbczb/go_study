package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
)

func GetInstance1() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Name: "Bard",
		}
	})
	return instance
}

func handleSingleton() {
	instance1 := GetInstance()
	instance2 := GetInstance()

	fmt.Println(instance1 == instance2) // true
}
