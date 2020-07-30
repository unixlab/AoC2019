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

	minDistance := 1000000
	for _, w1c := range wires[0].Coordinate {
		for _, w2c := range wires[1].Coordinate {
			if w1c.X == w2c.X && w1c.Y == w2c.Y {
				currentDistance := getDistance(0, 0, w1c.X, w1c.Y)
				if currentDistance < minDistance {
					minDistance = currentDistance
				}
			}
		}
	}
	fmt.Println(minDistance)
}
