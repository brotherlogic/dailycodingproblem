package main

import (
	"fmt"
	"time"
)

func solve1Brute(a []int, k int) bool {
	for i := range a {
		for j := range a[i:] {
			if a[i]+a[j] == k {
				return true
			}
		}
	}

	return false
}

func solve1OnePass(a []int, k int) bool {
	wantToSee := make([]bool, k)
	for i := range a {
		if wantToSee[a[i]] {
			return true
		}
		wantToSee[k-a[i]] = true
	}

	return false
}

func main() {
	iters := 10000

	t := time.Now()
	val := false
	for i := 0; i < iters; i++ {
		val = solve1Brute([]int{10, 15, 3, 7}, 17)
	}
	bruteTook := time.Now().Sub(t)
	fmt.Printf("%v (%v)\n", val, bruteTook)

	t = time.Now()
	val = false
	for i := 0; i < iters; i++ {
		val = solve1OnePass([]int{10, 15, 3, 7}, 17)
	}
	onePassTook := time.Now().Sub(t)
	fmt.Printf("%v (%v)\n", val, onePassTook)

}
