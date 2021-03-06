package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

//当多个连续的参数为同样的类型时，可以仅声明最后一个参数类型
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
