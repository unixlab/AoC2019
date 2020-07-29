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
	grid  [10][10]bool
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

func checkLos(y1 int, x1 int, y2 int, x2 int) bool {
	var xRatio, yRatio, xStep, yStep, xCur, yCur float64
	epsilon := 1e-5
	angle := getAngle(y1, x1, y2, x2)
	yRatio = angle / 90
	xRatio = (90 - angle) / 90
	xStep = 1 * xRatio
	yStep = 1 * yRatio

	done := getDistance(y1, x1, y2, x2)
	for {
		done -= xStep
		done -= yStep
		if done < epsilon {
			break
		}
		xCur += xStep
		yCur += yStep
		_, xFrac := math.Modf(xCur)
		_, yFrac := math.Modf(yCur)
		if (xFrac < epsilon || xFrac > 1.0-epsilon) && (yFrac < epsilon || yFrac > 1.0-epsilon) {
			var posY, posX int
			if y1 > y2 {
				posY = y1 - int(math.Round(yCur))
			} else {
				posY = y1 + int(math.Round(yCur))
			}
			if x1 > x2 {
				posX = x1 - int(math.Round(xCur))
			} else {
				posX = x1 + int(math.Round(xCur))
			}
			if grid[posY][posX] {
				return false
			}
		}
	}
	return true
}

func getDistance(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2))
}

func getAngle(y1 int, x1 int, y2 int, x2 int) float64 {
	a := int2floatAbs(y1 - y2)
	b := int2floatAbs(x1 - x2)
	c := math.Sqrt(a*a + b*b)
	alpha := math.Acos((b*b + c*c - a*a) / (2 * b * c))
	degree := alpha * 180 / math.Pi
	return degree
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
				for _, cord := range cords {
					if x == cord.X || y == cord.Y {
						continue
					}
					if checkLos(y, x, cord.Y, cord.X) {
						counter++
					}
				}
			}
			if counter > best {
				best = counter
			}
		}
	}

	fmt.Println(best)
}
