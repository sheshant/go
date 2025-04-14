package main

import "fmt"

func headRecursion(n int) int {
	// Base case
	if n <= 0 {
		return 0
	}
	// Recursive call
	return n + headRecursion(n-1)
}

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}

}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
