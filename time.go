package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	time.Sleep(10 * time.Second)
	t2 := time.Now()
	t3 := t2.Sub(t1).Minutes()
	fmt.Println(t3)

	fmt.Println(time.Now().After(time.Now().Add(-10 * time.Minute)))
}
