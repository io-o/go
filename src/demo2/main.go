package main

import "fmt"

func main() {
	n := 10
	// 判断 
	if n >= 20 {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}

	// for循环
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	// for循环的初始语句和结束语句都可以省略，例如：
	i := 0
	for i < 10 {
		fmt.Println("hh", i)
		i++
	} 
}