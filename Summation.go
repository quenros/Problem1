package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&n)

	//first method
	summation := sum_to_n_a(n)
	fmt.Printf("The summation from 1 to %d is: %d \n", n, summation)

	//second method
	summation = sum_to_n_b(n)
	fmt.Printf("The summation from 1 to %d is: %d \n", n, summation)

	//third method
	summation = sum_to_n_c(n)
	fmt.Printf("The summation from 1 to %d is: %d \n", n, summation)
}

func sum_to_n_a(n int) int {
	//Time complexity: O(1), since we are using arithmetics to calcuate the summation.
	//Most efficient method out of the other methods.
	return (n * (n + 1)) / 2
}

func sum_to_n_b(n int) int {
	//Time complexity is O(n) because we have a loop that iterates from 1 to n.
	//We might intuitively arrive at this solution, although it is not the best.
	summation := 0

	for i := 1; i <= n; i++ {
		summation += i
	}

	return summation
}

func sum_to_n_c(n int) int {
	//Time complexity is O(n) because we use recursion to call n times until it reaches 1.
	//I would not write it this way because if the code gets longer, it will get harder to read.

	if n == 1 {
		return 1
	}
	//eg: 5 + 4 + 3 + 2 + 1
	return n + sum_to_n_c(n-1)
}
