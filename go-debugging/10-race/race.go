package main

import (
	"fmt"
)

// Spot the issue?

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
