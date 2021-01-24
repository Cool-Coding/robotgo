package main

import (
	"os"
	"robotgo/server"
)

func main() {
	server.Run_Robot(os.Args[1])

	//arry := [10]int32{1,2,3,4,5,6,7,8,9,10}
	//
	//length := len(arry)
	//packLength := 2
	//startIndex := 0
	//for   {
	//	endIndex := startIndex + packLength
	//	if endIndex > length {
	//		endIndex = length
	//	}
	//
	//	a := arry[startIndex:endIndex]
	//
	//	println(a[0])
	//	println(a[1])
	//
	//	startIndex = endIndex
	//
	//	if startIndex >= length {
	//	break
	//	}
	//}
}
