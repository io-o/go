package main

import "fmt"

func main () {
	// var a [3]int
	// var b [4]int
	// fmt.Println(a)
	// fmt.Println(b)

	// 数组初始化
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// fmt.Println(cityArray)

	// 自动推导长度
	//var bArray = [...]bool{true, false, false, true, true}
	// fmt.Println(bArray)

	// 使用索引值
	//var numArray = [...]string{1: "golang", 7: "java"}
	// fmt.Println(numArray)

	// 数组循环
	// for i := 0; i < len(cityArray); i++{
	// 	fmt.Println(cityArray[i])
	// }

	// range循环
	// for i, v := range cityArray{
	// 	fmt.Println(i, v)
	// }

	// 只取 value值
	// for _, v := range cityArray{
	// 	fmt.Println(v)
	// }

	// 二维数组

	//  一个长度为3的数组 值是长度为2的数组
	// cityArray := [3][2]string{
	// 	{"北京", "天津"},
	// 	{"上海", "苏州"},
	// 	{"济南", "滕州"},
	// }
	// fmt.Println(cityArray)

	// 遍历二维数组
	// for _, v := range cityArray{
	// 	for _, v1 := range v {
	// 		fmt.Println(v1)
	// 	}
	// }

	// 数组求和
	a := [...]int{1, 3, 5, 7, 8}
	// num := 0
	// for _, v := range a{
	// 	num += v
	// }
	// fmt.Println(num)


}