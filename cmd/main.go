package main

import (
	"fmt"

	"github.com/mikepianka/ordinal"
)

func main() {
	for i := -1; i < 1025; i++ {
		fmt.Println(ordinal.Itoa(i))
	}
}
