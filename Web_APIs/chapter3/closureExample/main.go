package main

import (
	"fmt"
)

// This function returns another function
func generator() func() int { // Outer function
	var i = 0
	return func() int { // inner function
		i++
		return i
	}
}

func main() {
	numGenerator := generator()
	for i := 0; i < 5; i++ {
		fmt.Println(numGenerator(), "\t")
	}
}