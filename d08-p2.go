package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
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

		var finalPicture [25][6]int
		for y := 0; y < 6; y++ {
			for x := 0; x < 25; x++ {
				pixel := -1
				for _, picture := range pictures {
					if pixel == -1 && picture[x][y] != 2 {
						pixel = picture[x][y]
					}
				}
				finalPicture[x][y] = pixel
			}
		}

		for y := 0; y < 6; y++ {
			for x := 0; x < 25; x++ {
				if finalPicture[x][y] == 0 {
					color.New(color.BgBlack).Printf(" ")
				} else {
					color.New(color.BgWhite).Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
	}
}
