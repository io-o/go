package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true
	balance := 10000.0
	money := 0.0
	note := ""
	details := "收支\t账户\t收入\t说明"
	for{
		fmt.Println("\n`````````````家庭收支````````````````")
		fmt.Println("              1收支明细")
		fmt.Println("              2主持明细")
		fmt.Println("              3上手明细")
		fmt.Println("              4退出软件")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println("              1收支明细")
				fmt.Println(details)
			case "2":
				fmt.Println("本次收入金额:")
				fmt.Scanln(&money)
				balance += money
				fmt.Println("本次收入说明:")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance, money, note )
			case "3":
				fmt.Println("              3上手明细")
			case "4":
				loop = false
			default :
				fmt.Println("请输入正确")
					
		}
		if !loop {
			break
		}
	}
	fmt.Println("已退出")
}