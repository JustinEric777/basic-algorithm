package main

import "fmt"

func main() {
	a := -8
	b := -1
	fmt.Printf("%8b\n", a)
	fmt.Printf("%8b\n", b)
	fmt.Println(a & b)
}
