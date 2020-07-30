package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Wire struct {
	Coordinate []Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func main() {
	file, _ := os.Open("d03-input.txt")
	scanner := bufio.NewScanner(file)

	var wires []Wire

	for scanner.Scan() {
		var coordinates []Coordinate
		offsetX := 0
		offsetY := 0
		for _, move := range strings.Split(scanner.Text(), ",") {
			step, _ := strconv.Atoi(move[1:])
			if move[0:1] == "U" {
				for i := step; i > 0; i-- {
					offsetY += 1
					coordinates = append(coordinates, Coordinate{offsetX, offsetY})
				}
			} else if move[0:1] == "D" {
				for i := step; i > 0; i-- {
					offsetY -= 1
					coordinates = append(coordinates, Coordinate{offsetX, offsetY})
				}
			} else if move[0:1] == "R" {
				for i := step; i > 0; i-- {
					offsetX += 1
					coordinates = append(coordinates, Coordinate{offsetX, offsetY})
				}
			} else if move[0:1] == "L" {
				for i := step; i > 0; i-- {
					offsetX -= 1
					coordinates = append(coordinates, Coordinate{offsetX, offsetY})
				}
			} else {
				fmt.Printf("error on %s\n", move)
			}
		}
		wires = append(wires, Wire{coordinates})
	}

	minSteps := 1000000
	w1s := 0
	w2s := 0
	for _, w1c := range wires[0].Coordinate {
		w1s++
		w2s = 0
		for _, w2c := range wires[1].Coordinate {
			w2s++
			if w1c.X == w2c.X && w1c.Y == w2c.Y {
				if w1s+w2s < minSteps {
					minSteps = w1s + w2s
				}
			}
		}
	}
	fmt.Println(minSteps)
}
