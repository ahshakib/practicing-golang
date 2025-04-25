package main
import "fmt"

func main() {
	var s []int // declare a slice without initializing it
	fmt.Println(s) // nil slice
	fmt.Println("Length of slice s:", len(s)) // Length of slice s: 0
	fmt.Println("Capacity of slice s:", cap(s)) // Capacity of slice s: 0

	s = append(s, 1, 2, 3) // append elements to the slice
	fmt.Println(s) // [1 2 3]
	fmt.Println("Length of slice s:", len(s)) // Length of slice s: 3
	fmt.Println("Capacity of slice s:", cap(s)) // Capacity of slice s: 3
}