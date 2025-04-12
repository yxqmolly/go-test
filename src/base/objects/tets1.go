package main

import "fmt"



func main() {
	var person = &Person{
		Id: 1,
		Name: "nora",
	}
	// fmt.Println(person)
	fmt.Println(person.toString())
}
