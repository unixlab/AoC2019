package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getFuel(mass float64) int {
	return int(math.Floor(mass/3)) - 2
}

func main() {
	sum := 0
	file, _ := os.Open("d01-input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		sum += getFuel(mass)
	}
	fmt.Println(sum)
}
