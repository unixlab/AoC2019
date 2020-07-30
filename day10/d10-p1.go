package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	cords []cord
	grid  [25][25]bool
)

type cord struct {
	Y int
	X int
}

func int2floatAbs(number int) float64 {
	return math.Abs(float64(number))
}

func checkInRow(y int, x int) int {
	counter := 0
	for i := x + 1; i < len(grid); i++ {
		if grid[y][i] {
			counter++
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if grid[y][i] {
			counter++
			break
		}
	}
	return counter
}

func checkInColumn(y int, x int) int {
	counter := 0
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] {
			counter++
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if grid[i][x] {
			counter++
			break
		}
	}
	return counter
}

func getDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2))
}

func getAngle(y1 int, x1 int, y2 int, x2 int) float64 {
	degree := 180 - (180 / math.Pi * math.Atan2(float64(x1-x2), float64(y1-y2)))

	if x2 > x1 {
		degree += 90
	}
	if y2 > y1 {
		degree += 180
	}
	return degree
}

func appendIfNotExists(array []float64, element float64) []float64 {
	for _, v := range array {
		if v == element {
			return array
		}
	}
	return append(array, element)
}

func main() {
	file, _ := os.Open("d10-input.txt")
	scanner := bufio.NewScanner(file)

	x := 0
	y := len(grid)
	for scanner.Scan() {
		x = 0
		y--
		for _, char := range strings.Split(scanner.Text(), "") {
			if char == "#" {
				grid[y][x] = true
				cords = append(cords, cord{y, x})
			} else {
				grid[y][x] = false
			}
			x++
		}
	}

	best := 0

	for y, yV := range grid {
		for x, xV := range yV {
			counter := 0
			if xV {
				counter += checkInColumn(y, x)
				counter += checkInRow(y, x)
				var angles []float64
				for _, cord := range cords {
					if x == cord.X || y == cord.Y {
						continue
					}
					angle := getAngle(y, x, cord.Y, cord.X)
					angles = appendIfNotExists(angles, angle)
				}
				counter += len(angles)
			}
			if counter > best {
				best = counter
			}
		}
	}

	fmt.Println(best)
}
