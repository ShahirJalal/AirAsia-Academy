package main

import (
	"fmt"
	"go/constant"
)

func main() {
	counter := 18
	for i := 0; i < 10; i++ {
		fmt.Println("%v x %v = %v", counter, i, counter*i)
	}
}