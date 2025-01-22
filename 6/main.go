package main

import "fmt"

type Person struct {
	name string
	age int
}

type print interface {
	display()
}


func (p Person) display() {
	fmt.Printf("Name: %s | Age: %d\n", p.name, p.age)
}

func display(p print) {
	p.display()
}

func main() {
	var myMap map[string]int = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["two"] = 3
	fmt.Println(myMap)
	fmt.Println(len(myMap))

	var a Person = Person{"Shakib", 24}
	a.name = "Aftab"
	fmt.Println(a)
	display(a)

	var struct2 = struct {
		name string
		age int
	} {"Shakib", 24}
	fmt.Println(struct2)
}