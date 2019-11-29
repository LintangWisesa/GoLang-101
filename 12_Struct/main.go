package main

import (
	"fmt"
	"strconv"
)

// 1. declare a person struct
type person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// 2. Hello method (value receiver)
func (x person) hello() string {
	return "Hello, " + x.firstName + " " + x.lastName + " " + strconv.Itoa(x.age) + "th"
}

// 3. has birthday method (pointer receiver)
// param pointer receiver should be same as before: x
func (x *person) hasBirthday() {
	x.age++
}

func main() {
	// 1. call struct
	personA := person{firstName: "Andi", lastName: "Wijaya", city: "Jakarta", gender: "Pria", age: 28}
	fmt.Println(personA)
	fmt.Println(personA.age)

	personB := person{"Budi", "Susilo", "Bandung", "Pria", 27}
	fmt.Println(personB)
	fmt.Println(personB.city)

	// 2. call hello method
	fmt.Println(personA.hello())

	// 3. call hasBirthday method
	personA.hasBirthday()
	personA.hasBirthday()
	fmt.Println(personA.hello())
}
