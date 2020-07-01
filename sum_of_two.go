package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

//first string - target number
//second string - n-numbers
//return 1 if a+b=X exists
func main() {

	file, _ := os.Create("problems/output.txt")
	defer file.Close()

	f, _ := os.Open("problems/input.txt")
	defer f.Close()

	set := make(map[int]bool)

	var s scanner.Scanner
	s.Init(bufio.NewReader(f))
	s.Mode = scanner.ScanInts

	s.Scan()
	target, _ := strconv.Atoi(s.TokenText())

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {

		val, _ := strconv.Atoi(s.TokenText())
		fmt.Println(val)
		if set[target-val] {
			file.WriteString("1")
			return
		}
		set[val] = true
	}
	file.WriteString("0")
}
