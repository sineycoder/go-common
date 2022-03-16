package main

import (
	"fmt"

	"github.com/go-common/common_tool"
)

/**
 @author: nizhenxian
 @date: 2022/3/16 10:03:16
**/
func main() {
	arr := []float32{1, 2.2, 1, 2}
	common_tool.Sort[float32](arr, func(i, j float32) bool {
		return i < j
	})
	fmt.Println(arr)
}
