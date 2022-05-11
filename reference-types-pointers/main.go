package main

import (
	"fmt"
)

// reference types (pointers, slices, maps, functions, channels)

// interface type

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Println("A %s says %s", a.Name, a.Sound)
}

func main() {

	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)

	*myFirstPointer = 15

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call, x is now", x)

	// Slices
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "cow")
	animals = append(animals, "wolf")

	fmt.Println(animals)

	for _, x := range animals {
		fmt.Println(x)
	}

	fmt.Println("First two elements are: ", animals[0:2])

	// Maps
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "four")

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}

	// Functions
	z := addTwoNumbers(2, 4)
	fmt.Println(z)

	f := addTwoNumber(5, 6)
	fmt.Println(f)

	myTotal := sumMany(5, 6, 7, 8)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "Ben"
	dog.Sound = "woof"
	dog.Says()

}

func changeValueOfPointer(num *int) {
	*num = 25
}

func addTwoNumbers(x, y int) int {
	return x + y
}

func addTwoNumber(x, y int) (sum int) {
	sum = x + y
	return
}

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}
