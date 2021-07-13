package main

import (
	"fmt"
)

func main() {
	ret := add(0, 2)
	fmt.Println("程序运行结束", ret)
}

func add(i, j int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if i < 0 {
		panic("i值不能小于0")
	}
	i = j / i
	return 1 + j
}
