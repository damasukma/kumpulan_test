package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		data := fibonacci(i)
		fmt.Println(data)
	}
}

func fibonacci(n int) int {

	if n > 1 {
		return fibonacci(n-1) + fibonacci(n-2)

	}
	return n

}
