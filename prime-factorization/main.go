package main

import "fmt"

func prime_factors(n int) []int {
	factors := make([]int, 0)
	i := 2
	for i*i <= n {
		r := n % i
		if r != 0 {
			i += 1
		} else if r == 0 {
			n /= i
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func main() {
	fmt.Print(prime_factors(1200)) // [2 2 2 2 3 5 5]
}
