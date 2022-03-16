package common_tool

import "sort"

/**
 @author: nizhenxian
 @date: 2022/3/13 15:08:41
**/
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sort[T Ordered](arr []T, f func(i, j T) bool) {
	sort.Slice(arr, func(i, j int) bool {
		return f(arr[i], arr[j])
	})
}
