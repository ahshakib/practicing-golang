package main
import "fmt"
func main() {
	var x, t int
	fmt.Scan(&x, &t)
	if t >= x {
		fmt.Println(0)
	} else {
		fmt.Println(x - t)
	}
}