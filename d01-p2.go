package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func main() {
	var fuelForModules, fuelForFuel, fuelSum, fuelSumOverall float64
	file, _ := os.Open("d01-input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// reset variables
		fuelForModules = 0
		fuelForFuel = 0
		fuelSum = 0

		// get fuel for modules
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		fuelForModules += getFuel(mass)

		// get inital fuel for fuel
		fuelForFuel = getFuel(fuelForModules)
		fuelSum = fuelForModules + fuelForFuel

		// loop as long as fuel for fuel is < 0
		for getFuel(fuelForFuel) >= 0 {
			fuelForFuel = getFuel(fuelForFuel)
			fuelSum += fuelForFuel
		}

		// sum for current module
		fuelSumOverall += fuelSum
	}
	fmt.Printf("%0.0f\n", fuelSumOverall)
}
