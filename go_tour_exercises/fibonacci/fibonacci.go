package main

import "fmt"

// fibonacci function returns a closure that generates successive Fibonacci numbers.
func fibonacci() func() int {
	// Two variables to keep track of the current and previous numbers in the Fibonacci sequence.
	a, b := 0, 1

	// The closure returned by the function.
	// This closure captures and updates the variables a and b.
	return func() int {
		// Store the current value of a, which is the next number in the sequence.
		ret := a

		// Update a and b to the next values in the Fibonacci sequence.
		a, b = b, a+b

		// Return the next Fibonacci number.
		return ret
	}
}

func main() {
	// Get the Fibonacci closure function.
	f := fibonacci()

	// Print the first 10 Fibonacci numbers.
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
