package main

import (
	"encoding/binary"
	"fmt"
)

//read 2 numbers from cli
func main()  {
	var a,b int
	fmt.Scanf("%d %d", &a, &b)
	var pa = make([]byte, 64)
	binary.LittleEndian.PutUint32(pa, uint32(a))
	pb := byte(b)
	fmt.Println(pa, pb)
}
