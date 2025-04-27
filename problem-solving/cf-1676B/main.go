package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		minVal := 100000000
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
			if minVal > a[i] {
				minVal = a[i]
			}
		}
		total := 0
		for i := 0; i < n; i++ {
			total += (a[i] - minVal) // <== no if condition
		}
		fmt.Println(total)
	}
}
