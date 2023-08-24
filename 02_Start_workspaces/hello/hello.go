package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))

	// Use func Int defined by us.
	fmt.Println((reverse.Int(232136)))
}
