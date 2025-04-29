package main

import (
	"fmt"
	"sort"
)

func main() {
	var t, n, i int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&n)
		a := make([]int, n)
		for i = 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		if n == 1 {
			fmt.Println("YES")
			continue
		}
		sort.Ints(a)
		f := false
		for i = 0; i < n-1; i++ {
			if a[i+1] - a[i] > 1 {
				fmt.Println("NO")
				f = true
				break
			}
		}
		if !f {
			fmt.Println("YES")
		}
	}
}