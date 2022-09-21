package main

import "fmt"

// import "fmt"

func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}

func main() {
	Greet("John Doe")
}