package main
import (
	"fmt"
	"errors"
)

func main() {
	greetings()
	var a, b int = getUserInput()
	var sum int = add(a, b)
	fmt.Println(sum)
	var div, mod, err = getDivision(a, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Division: %d, Modules: %d\n", div, mod)
}

func add(a int, b int) int {
	return a + b
}

func getUserInput() (int, int) {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	return a, b
}

func getDivision(a int, b int) (int, int, error) {
	var err error
	if(b == 0) {
		err = errors.New("division by zero is not allowed")
		return 0, 0, err
	}
	return a/b, a%b, err
}

func greetings () {
	fmt.Println("Welcom to Go Programming")
}