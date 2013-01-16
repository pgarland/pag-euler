package main

import (
	"fmt"
	"os"
	"strconv"
)

func lower_floor_div(x, y int) int {
	floor := x/y
	if x % y == 0 {
		floor = floor - 1
	}
    return floor
}

func sum_multiples_less_than (factor, y int) int {
    greatest_factor := lower_floor_div(y, factor)
    return (factor * ((greatest_factor + 1) * greatest_factor)/2)
}

func soln (factors []int, top int) int {
	n_factors := len(factors)
	// Create a buffered channel to push the results into
	ci := make(chan int, n_factors)

	// Spin off a goroutine for each factor
	for i:=0; i < n_factors; i++ {
		factor := factors[i]
		go func() {
			ci <- sum_multiples_less_than(factor, top)
		}()
	}

	// Collect the results of the goroutines
	sum := 0
	for i := 0; i < n_factors; i++ {
		sum += <- ci
	}

	return sum
}

func main () {
	n_factors := len(os.Args) - 2
	factors := make( []int, n_factors)

	top, err := strconv.Atoi(os.Args[n_factors + 1])
	for i:=0; i < n_factors; i++ {
		factors[i], err = strconv.Atoi(os.Args[i + 1])
		if err != nil {
			fmt.Println(err)
		}
	}
	sum := soln(factors, top)
	fmt.Println(sum)
}
