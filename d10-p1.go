package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}
	return result
}

func main() {
	file, _ := os.Open("d10-input.txt")
	scanner := bufio.NewScanner(file)

	var positions [][2]float64

	x := 0
	for scanner.Scan() {
		for y, char := range strings.Split(scanner.Text(), "") {
			if char == "#" {
				positions = append(positions, [2]float64{float64(x), float64(y)})
			}
		}
		x++
	}

	for positionIndex, positionValue := range positions {
		for currentIndex, currentValue := range positions {
			if positionIndex == currentIndex {
				continue
			}
			fmt.Printf("%0.f.%0.f to %0.f.%0.f - ", positionValue[0], positionValue[1], currentValue[0], currentValue[1])
			wayOnX := math.Abs(currentValue[0] - positionValue[0])
			wayOnY := math.Abs(currentValue[1] - positionValue[1])

			if wayOnY == 0 || wayOnX == 0 {
				fmt.Printf("skipped\n")
				continue
			}

			greatestCommonDivisor := greatestCommonDivisor(int(wayOnX), int(wayOnY))
			if greatestCommonDivisor == 1 {
				fmt.Printf("not possible\n")
				continue
			}

			if wayOnY == wayOnX {
				var x, y float64
				if currentValue[0] - positionValue[0] > 0 && currentValue[1] - positionValue[1] > 0 {
					x = currentValue[0] - 1
					y = currentValue[1] - 1
				}
				if currentValue[0] - positionValue[0] < 0 && currentValue[1] - positionValue[1] > 0 {
					x = currentValue[0] + 1
					y = currentValue[1] - 1
				}
				if currentValue[0] - positionValue[0] > 0 && currentValue[1] - positionValue[1] < 0 {
					x = currentValue[0] - 1
					y = currentValue[1] + 1
				}
				if currentValue[0] - positionValue[0] < 0 && currentValue[1] - positionValue[1] < 0 {
					x = currentValue[0] + 1
					y = currentValue[1] + 1
				}
				for _, v := range positions {
					if v[0] == x && v[1] == y {
						fmt.Printf("hit on %0.f.%0.f", x, y)
					}
				}
				fmt.Println()
				continue
			}

			fmt.Printf("x=%0.f y=%0.f r=%d\n", wayOnX, wayOnY, greatestCommonDivisor)

		}
		break
	}
}
