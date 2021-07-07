package main

import "fmt"

func main() {
	i := 2
	var msg string
	switch i {
	case 1, 2, 3:
		msg = "等于1/2/3"
	case 4:
		msg = "等于4"
	default:
		msg = "整了个默认值"
	}
	fmt.Println(msg)
	switch i {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		msg = "等于1/2/3"
	case 4:
		msg = "等于4"
	default:
		msg = "整了个默认值"
	}
	fmt.Println(msg)
}
