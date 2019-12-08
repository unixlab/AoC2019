package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("d08-input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pixel := strings.Split(scanner.Text(), "")
		var pictures [][25][6]int
		var picture [25][6]int
		x := 0
		y := 0
		for _, v := range pixel {
			pix, _ := strconv.Atoi(v)
			picture[x][y] = pix

			x++

			if x == 25 {
				x = 0
				y++
			}
			if y == 6 {
				y = 0
				pictures = append(pictures, picture)
			}
		}

		var min0, max1, max2 int
		for _, picture := range pictures {
			var numbers [3]int
			for _, x := range picture {
				for _, y := range x {
					numbers[y]++
				}
			}
			if numbers[0] < min0 || min0 == 0 {
				min0 = numbers[0]
				max1 = numbers[1]
				max2 = numbers[2]
			}
		}
		fmt.Printf("0=%d 1=%d 2=%d 1x2=%d\n", min0, max1, max2, max1*max2)
	}
}
