package main
import (
	"fmt"
)

func main() {
	var t, n, i int
	fmt.Scan(&t)
	for t > 0 {
		fmt.Scan(&n)
		a := make([]int, n)
		b := make([]int, n)
		for i = 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		for i = 0; i < n; i++ {
			fmt.Scan(&b[i])
		}
		max := a[n-1]
		for i = 0; i < n-1; i++ {
			if a[i] - b[i+1] > 0 {
				max += a[i] - b[i+1]
			}
		}
		fmt.Println(max)
		t--
	}
}