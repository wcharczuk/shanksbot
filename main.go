package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "please provide an input prime")
		os.Exit(1)
	}

	rawPrime := os.Args[1]
	prime, err := strconv.ParseInt(rawPrime, 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid integer prime provided;", err.Error())
		os.Exit(1)
	}
	if prime < 0 {
		fmt.Fprintln(os.Stderr, "please provide a positive prime")
		os.Exit(1)
	}
	if !isPrime(prime) {
		fmt.Fprintln(os.Stderr, "please provide a valid prime, i.e. a number divisible by itself and 1 exclusively")
		os.Exit(1)
	}

	output := primeReciprocalRepeatsAfter(int(prime))
	fmt.Println("input:", prime, "output:", output)
}

func isPrime(prime int64) bool {
	return big.NewInt(prime).ProbablyPrime(0)
}

func primeReciprocalRepeatsAfter(prime int) int {
	seen := make([]int, 0, prime-1)

	dividend := 0
	remainder := 1
	for !repeats(seen) {
		var zeros int
		for remainder < prime {
			remainder = remainder * 10
			zeros++
		}

		dividend = remainder / prime
		remainder = remainder % prime
		if zeros > 1 {
			for x := 0; x < zeros-1; x++ {
				seen = append(seen, 0)
			}
			seen = append(seen, dividend)
		} else {
			seen = append(seen, dividend)
		}
	}
	return len(seen) >> 1
}

func repeats(i []int) bool {
	ilen := len(i)
	if ilen < 2 {
		return false
	}
	if ilen%2 != 0 {
		return false
	}
	middle := ilen >> 1

	for x := 0; x < middle; x++ {
		if i[x] != i[x+middle] {
			return false
		}
	}
	return true
}
