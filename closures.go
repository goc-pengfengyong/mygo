package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	/**
	我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。
	这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值。
	*/

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
