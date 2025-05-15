package main
import "fmt"

var c = 0
var dp [100005]int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lis(n int, a []int) int {
	c++
	if dp[n] != -1 {
		return dp[n]
	}
	ans := 1
	for i := 0; i < n; i++ {
		if a[n] > a[i] {
			ans = max(ans, lis(i, a) + 1)
		}
	}
	dp[n] = ans
	return dp[n]
}

func main() {
	for i := range dp {
		dp[i] = -1
	}
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, lis(i, a))
	}
	fmt.Println(ans)
	fmt.Println("lis calls:", c)
}