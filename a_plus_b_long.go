package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitConvert(str string) []int64 {
	splitStr := strings.Split(str, "")
	arr := make([]int64, 0)

	for i := len(splitStr) - 1; i >= 0; i-- {
		tmp, _ := strconv.ParseInt(splitStr[i], 10, 64)
		arr = append(arr, tmp)
	}
	return arr
}
func maxLength(lenA, lenB int) int {
	if lenA > lenB {
		return lenA
	}
	return lenB
}

//read 2 strings from cli
func main() {

	var strA, strB string
	fmt.Scanf("%s", &strA)
	fmt.Scanf("%s", &strB)

	a := splitConvert(strA)
	b := splitConvert(strB)
	result := make([]int64, 0)
	maxLength := maxLength(len(a), len(b))

	var carry int64
	for i := 0; i < maxLength ; i++ {
		if i < len(a) && i < len(b) {
			result = append(result, (a[i]+b[i]+carry)%10)
			carry = (a[i] + b[i] + carry) / 10
			continue
		}
		if i < len(a) {
			result = append(result, (a[i]+carry)%10)
			carry = (a[i] + carry) / 10
		}
		if i < len(b) {
			result = append(result, (b[i]+carry)%10)
			carry = (b[i] + carry) / 10
		}
	}
	if carry > 0 {
		result = append(result, 1)
	}

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}
