package main

import "fmt"

func main() {
	slice1 := []int{1,2,3,4}
	slice1 = filter(func(n int) bool {
		return n % 2 == 0
	}, slice1)

	slice2 := []int{2,4,6,8}
	fmt.Printf("%p\n",&slice2)
	filterPtr(func(n *int) bool {
		return *n < 4
	}, &slice2)
	fmt.Printf("%p\n",&slice2)

	fmt.Println(slice1,slice2)
}

func filter[T any](f func (T) bool, s []T) []T {
	ret := make([]T,0,len(s))
	for _, v := range s {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func filterPtr[T any](f func (*T) bool, s *[]T) { 
	ret := make([]T,0,len(*s))
	for _, v := range *s {
		if f(&v) {
			ret = append(ret, v)
		}
	}
	*s = ret
	//so you can't do s = &ret because s is a copy of a pointer, it points to the same address
	//but it is a local copy
}
