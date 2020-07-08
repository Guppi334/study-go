package main

import (
	"fmt"

	"./mylib"
	"./mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()
	person := mylib.Person{Name: "Mike", Age: 24}
	fmt.Println(person)
}
