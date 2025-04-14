package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func main() {
	a := Person{"Alice", 30}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
