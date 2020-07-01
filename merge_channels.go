package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Merge2Channel(f func(int) int, in1 <-chan int, in2 <-chan int, out chan int, n int) chan int { //out chan<- int

	receiveByPair := func(chs ...<-chan int) (rs []int, ok bool) {
		ok = true
		for _, ch := range chs {
			r, ok2 := <-ch
			rs, ok = append(rs, r), ok && ok2
		}
		return rs, ok
	}
	pairChan := make(chan []int, n)
	go func() {
		defer close(pairChan)
		for pair, ok := receiveByPair(in1, in2); ok; pair, ok = receiveByPair(in1, in2) {
			pairChan <- pair
		}
	}()
	go func() {
		defer close(out)
		for rs := range pairChan {
			a, b := rs[0], rs[1]
			out <- f(a)+f(b)
		}
	}()
	return out
}

func someFunc(a int) int {
	return a
}
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func main() {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	out := make(chan int)

	kk:=Merge2Channel(someFunc, a, b, out, 20)
	fmt.Println()

	for v := range kk {
		fmt.Println(v)
	}
}
