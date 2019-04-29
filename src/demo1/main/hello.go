package main

import "fmt" // 必须用双引号 单引号报错

// 全局变量

var n int = 100

// 常量 用 const关键字 声明常量必须赋值
const a = "a"
const (
	b = "b"
	c = "c"
)


func main() {

	// 变量 批量声明
	var (
		name string
		age  int
		isOk bool
	)

	//  变量赋值
	var name1 string = "哈哈哈"

	// 多个赋值 逗号隔开
	var name2, age2 = "ggg", 26

	// 短变量声明 也是常用的声明方式
	name3 := "go"
	age3 := 100



	fmt.Println(n, name, age, isOk, name1, name2, age2, name3, age3) // 必须用双引号 单引号报错
}
