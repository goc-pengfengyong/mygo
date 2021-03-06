package main

import "fmt"

func main() {

	//创建一个空 map，需要使用内建函数 make：make(map[key-type]val-type)。
	m := make(map[string]int)
	//var n map[string]int

	fmt.Printf("%T", m)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	/**
	当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，
	该值表明了 map 中是否存在这个键。 这可以用来消除 键不存在
	和 键的值为零值 产生的歧义， 例如 0 和 ""。这里我们不需要值，
	所以用 空白标识符(blank identifier) _ 将其忽略。
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
