package main

import (
	"fmt"
)


func fibonacci(a, b, k, sum int) int{
	if k == 0 {
		return sum
	}
	sum += b
	k--
	return fibonacci(b, a + b, k, sum)
}

func main() {
	var k int
	fmt.Scanf("%d", &k)
	fmt.Println(fibonacci(0, 1, k, 0))
}
