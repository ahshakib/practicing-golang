package main
import (
	"fmt"
	"math"
)
var dp [100001]int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func jump(i int, h []int) int {
	if i == 0 {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}
	cost := math.MaxInt32
	
	a := jump(i-1, h) + abs(h[i] - h[i-1])
	if cost > a {
		cost = a
	}
	if i > 1 {
		b := jump(i-2, h) + abs(h[i] - h[i-2])
		if cost > b {
			cost = b
		}
	}
	dp[i] = cost
	return dp[i]
}

func main() {
	for i := range dp {
		dp[i] = -1
	}
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	fmt.Println(jump(n-1, h))
}