package main

import "fmt"

func main() {
	var p *int = new(int)
	var q int = 43
	p = &q
	fmt.Println(*p, q)
	*p = 100
	fmt.Println(*p, q)

	var slice = []int {1, 2, 3, 4, 5} // slice is a reference type
	var slice2 = slice
	fmt.Println(slice, slice2)
	slice[0] = 44
	fmt.Println(slice, slice2)

	//-------
	var arr = [5] int {1, 2, 3, 4, 5} // array is a value type
	fmt.Println(makeSquare(&arr))
}

func makeSquare(arr *[5]int) [5]int{
	for i := 0; i < 5; i++ {
		arr[i] = arr[i] * arr[i]
	}
	return *arr
}