package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}

	fmt.Println(3 & 1)
}
