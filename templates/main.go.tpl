package main

import (
	"fmt"
)

// Hallo is my hallo
func Hallo() string {
	return "Hello, world."

}

func main() {
	fmt.Printf("%s", Hallo())
}
