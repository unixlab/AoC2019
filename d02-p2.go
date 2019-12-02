package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringArrayToIntArray(stringNumbers []string) []int {
	var intNumbers []int
	for _, v := range stringNumbers {
		intNumber, _ := strconv.Atoi(v)
		intNumbers = append(intNumbers, intNumber)
	}
	return intNumbers
}

func runIntcode(noun int, verb int) int {
	file, _ := os.Open("d02-input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		offset := 0
		intCode := stringArrayToIntArray(strings.Split(scanner.Text(), ","))
		intCode[1] = noun
		intCode[2] = verb
		for intCode[offset] != 99 {
			if intCode[offset] == 1 {
				firstNumber := intCode[intCode[offset+1]]
				secondNumber := intCode[intCode[offset+2]]
				intCode[intCode[offset+3]] = firstNumber + secondNumber
			} else if intCode[offset] == 2 {
				firstNumber := intCode[intCode[offset+1]]
				secondNumber := intCode[intCode[offset+2]]
				intCode[intCode[offset+3]] = firstNumber * secondNumber
			} else {
				fmt.Println("error")
			}
			offset += 4
		}
		return (intCode[0])
	}
	return 0
}

func main() {
	for i := 0; i <= 100; i++ {
		for n := 0; n <= 100; n++ {
			if runIntcode(i, n) == 19690720 {
				fmt.Printf("noun=%d verb=%d answer=%d\n", i, n, 100*i+n)
			}
		}
	}
}
