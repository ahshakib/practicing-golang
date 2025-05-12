package main
import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		temp := a
		a = b
		b = temp % b
	}
	return a
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(gcd(a, b))
}