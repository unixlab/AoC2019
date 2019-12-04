package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkTwoAdjacentDigits(digit int) bool {
	stringDigit := strconv.Itoa(digit)
	stringDigitArray := strings.Split(stringDigit, "")

	for i := 0; i < len(stringDigitArray)-1; i++ {
		if stringDigitArray[i] == stringDigitArray[i+1] {
			return true
		}
	}

	return false
}

func checkNeverDecreasingDigits(digit int) bool {
	stringDigit := strconv.Itoa(digit)
	stringDigitArray := strings.Split(stringDigit, "")

	for i := 0; i < len(stringDigitArray)-1; i++ {
		firstNumber, _ := strconv.Atoi(stringDigitArray[i])
		secondNumber, _ := strconv.Atoi(stringDigitArray[i+1])
		if secondNumber < firstNumber {
			return false
		}
	}

	return true
}

func checkEvenNumbersOfDigits(digit int) bool {
	stringDigit := strconv.Itoa(digit)
	stringDigitArray := strings.Split(stringDigit, "")

	for _, v := range stringDigitArray {
		digitCounter := 0
		for i := 0; i < len(stringDigitArray); i++ {
			if v == stringDigitArray[i] {
				digitCounter++
			}
		}
		if digitCounter == 2 {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("d04-input.txt")
	scanner := bufio.NewScanner(file)

	lowerBoundary := 0
	upperBoundary := 0

	for scanner.Scan() {
		stringBoundary := strings.Split(scanner.Text(), "-")
		lowerBoundary, _ = strconv.Atoi(stringBoundary[0])
		upperBoundary, _ = strconv.Atoi(stringBoundary[1])
	}

	possiblePasswords := 0

	for lowerBoundary < upperBoundary {
		if checkTwoAdjacentDigits(lowerBoundary) && checkNeverDecreasingDigits(lowerBoundary) && checkEvenNumbersOfDigits(lowerBoundary) {
			possiblePasswords++
		}
		lowerBoundary++
	}

	fmt.Println(possiblePasswords)
}
