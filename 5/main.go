// Codeforces-2055A : Two Frogs
//https://codeforces.com/contest/2055/problem/A

package main

import (
	"fmt"
	"math"
)

func main() {
	var t, n, a, b, x int
	fmt.Scan(&t)
	for t > 0 {
		fmt.Scan(&n, &a, &b)
		x = int(math.Abs(float64(a) - float64(b)))
		if x == 2 || x % 2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}