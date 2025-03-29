package main
import "fmt"

type User struct {
	Name string
	Age int
	Salary float64
}

func main() {
	obj1 := User{
		Name: "Shakib",
		Age: 25,
		Salary: 1000.50,
	}
	p := &obj1 // p is a pointer to obj1, it stores the address of obj1 in memory
	// &obj1 is the address of obj1 in memory
	fmt.Println(*p) // dereference p to get the value of obj1
	fmt.Println(p.Name)
}