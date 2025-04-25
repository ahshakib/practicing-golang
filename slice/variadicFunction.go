package main
import "fmt"

func print(numbers ...int) { // variadic function & numbers is a slice of int
	fmt.Println("Numbers:", numbers)
	fmt.Println("Length of numbers:", len(numbers))
	fmt.Println("Capacity of numbers:", cap(numbers))
}

func main() {
	print(1, 2, 3) // passing variadic arguments
}