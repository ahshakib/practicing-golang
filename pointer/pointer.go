// pointer => address of memory(ram)
package main
import (
	"fmt"
)

// //pass by value
// func print(numbers [3]int) {
// 	fmt.Println(numbers)
// }

//pass by reference
func print(numbers *[3]int) {
	fmt.Println(numbers)
}

func main() {
	// x := 20
	// p := &x // p is a pointer to x, it stores the address of x in memory
	// // &x is the address of x in memory
	// *p = 30
	// fmt.Println("Value of x:", x) 
	// fmt.Println("Address of x:", p) // address of x in memory
	// fmt.Println("Value of x:", *p) // dereference p to get the value of x

	arr := [3]int{1,2,3}
	print(&arr)
}