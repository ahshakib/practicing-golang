package main
import "fmt"

var c int = 0
var dp [100001]int

func fib(n int) int {
	c++
	if n <= 1 {
		return n
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = fib(n-1) + fib(n-2)
	return dp[n]
}

func main() {
	for i := range dp {
		dp[i] = -1
	}
	var n int
	fmt.Scan(&n)
	fmt.Println(fib(n))
	fmt.Println("Number of calls:", c)
}