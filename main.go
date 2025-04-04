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
	fmt.Println("Choose a number > 0 and < 18,446,744,073,709,551,615")
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

	return true
}
