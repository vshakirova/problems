package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//return number without pair
func main() {
	file, _ := os.Open("problems/input-201.txt")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	set := make(map[int64]bool)

	for scanner.Scan() {
		str := scanner.Text()
		elem, _ := strconv.ParseInt(str, 10, 64)
		if set[elem] == true {
			set[elem] = false
			continue
		}
		set[elem] = true
	}
	file.Close()

	for k, v := range set {
		if v == true {
			fmt.Println(k)
			break
		}
	}
}
