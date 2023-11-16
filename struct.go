package main

import "fmt"

type Subscriber struct {
	name string
	age int
}

func main() {
	sub := Subscriber{"anya", 6}
	fmt.Println(sub.name)
	fmt.Println(sub.age)
}
