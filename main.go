package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Is it prime?!")
	fmt.Println("Choose a number > 2 and < 18,446,744,073,709,551,615")
	fmt.Println("---------------------")
	for {
		// input
		fmt.Print("Enter a number: ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		// validate
		num, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			fmt.Println("Not a valid input, womp womp..")
			fmt.Println(err.Error())
			fmt.Println()
			continue
		}

		// assemble the jury

		// well, is it?
		isPrime := isItPrime(num)
		if isPrime {
			fmt.Println("It's prime!")
		} else {
			fmt.Println("Not prime :(")
		}
		fmt.Println()
	}
}

func isItPrime(num uint64) bool {
	if num%2 == 0 {
		return false
	}

	// optimal witnesses
	var witnesses []int
	switch {
	case num < 2047:
		witnesses = []int{2}
	case num < 1373653:
		witnesses = []int{2, 3}
	case num < 9080191:
		witnesses = []int{31, 73}
	case num < 25326001:
		witnesses = []int{2, 3, 5}
	case num < 3215031751:
		witnesses = []int{2, 3, 5, 7}
	case num < 4759123141:
		witnesses = []int{2, 7, 61}
	case num < 1122004669633:
		witnesses = []int{2, 13, 23, 1662803}
	case num < 2152302898747:
		witnesses = []int{2, 3, 5, 7, 11}
	case num < 3474749660383:
		witnesses = []int{2, 3, 5, 7, 11, 13}
	case num < 341550071728321:
		witnesses = []int{2, 3, 5, 7, 11, 13, 17}
	case num < 3825123056546413051:
		witnesses = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	case num <= 18446744073709551615:
		witnesses = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	default:
		witnesses = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	}
	fmt.Println(witnesses)

	// find m and d for n^m*d+1 = num
	m := 0
	d := num - 1
	for d%2 == 0 {
		d = d / 2
		m++
	}
	fmt.Printf("2^%d*%d + 1\n", m, d)

	// a^d = 1%n - for each witness
	isPrime := true
	for _, wit := range witnesses {

		// optimization so we dont have to calculate a^d
		runningTotal := uint64(1)
		for i := uint64(0); i < d; i++ {
			// runningTotal = (runningTotal * uint64(wit)) % num
			runningTotal = runningTotal * uint64(wit)
		}
		// fmt.Printf("runningTotal: %d^%d = %d\n", wit, d, runningTotal)

		// ******* we may need to do some checks with other values of 2^m and d **********

		runningTotal = runningTotal % num
		fmt.Printf("%d^%d mod %d = %d\n", wit, d, num, runningTotal)

		if runningTotal != 1 || runningTotal != num-1 {
			// test := (uint64(wit) ^ d) % num
			// if test != 1 || test != num-1 {
			isPrime = false
			break
		}
	}

	return isPrime
}
