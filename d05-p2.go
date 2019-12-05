package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringArrayToIntArray(stringNumbers []string) []int {
	var intNumbers []int
	for _, v := range stringNumbers {
		intNumber, _ := strconv.Atoi(v)
		intNumbers = append(intNumbers, intNumber)
	}
	return intNumbers
}

func parseOptCode(instructionCode int) (int, []int) {
	instructionString := strconv.Itoa(instructionCode)
	instructionLength := len(instructionString)
	opcode := 0
	if instructionLength > 1 {
		opcode, _ = strconv.Atoi(instructionString[instructionLength-2:])
	} else {
		return instructionCode, []int{0, 0, 0}
	}
	instructionOffset := instructionLength - 2
	parameterModes := []string{"0", "0", "0"}
	for i := instructionOffset; i > 0; i-- {
		parameterModes[instructionOffset-i] = instructionString[i-1 : i]
	}
	return opcode, stringArrayToIntArray(parameterModes)
}

func intCodeComputer(intCode []int, input int) {
	offset := 0
	for {
		opcode, modes := parseOptCode(intCode[offset])
		if opcode == 99 {
			break
		} else if opcode == 1 {
			var firstNumber, secondNumber int
			if modes[0] == 0 {
				firstNumber = intCode[intCode[offset+1]]
			} else {
				firstNumber = intCode[offset+1]
			}
			if modes[1] == 0 {
				secondNumber = intCode[intCode[offset+2]]
			} else {
				secondNumber = intCode[offset+2]
			}
			intCode[intCode[offset+3]] = firstNumber + secondNumber
			offset += 4
		} else if opcode == 2 {
			var firstNumber, secondNumber int
			if modes[0] == 0 {
				firstNumber = intCode[intCode[offset+1]]
			} else {
				firstNumber = intCode[offset+1]
			}
			if modes[1] == 0 {
				secondNumber = intCode[intCode[offset+2]]
			} else {
				secondNumber = intCode[offset+2]
			}
			intCode[intCode[offset+3]] = firstNumber * secondNumber
			offset += 4
		} else if opcode == 3 {
			intCode[intCode[offset+1]] = input
			offset += 2
		} else if opcode == 4 {
			if modes[0] == 0 {
				fmt.Println(intCode[intCode[offset+1]])
			} else {
				fmt.Println(intCode[offset+1])
			}
			offset += 2
		} else if opcode == 5 {
			if modes[0] == 0 {
				if intCode[intCode[offset+1]] != 0 {
					if modes[1] == 0 {
						offset = intCode[intCode[offset+2]]
					} else {
						offset = intCode[offset+2]
					}
				} else {
					offset += 3
				}
			} else {
				if intCode[offset+1] != 0 {
					if modes[1] == 0 {
						offset = intCode[intCode[offset+2]]
					} else {
						offset = intCode[offset+2]
					}
				} else {
					offset += 3
				}
			}
		} else if opcode == 6 {
			if modes[0] == 0 {
				if intCode[intCode[offset+1]] == 0 {
					if modes[1] == 0 {
						offset = intCode[intCode[offset+2]]
					} else {
						offset = intCode[offset+2]
					}
				} else {
					offset += 3
				}
			} else {
				if intCode[offset+1] == 0 {
					if modes[1] == 0 {
						offset = intCode[intCode[offset+2]]
					} else {
						offset = intCode[offset+2]
					}
				} else {
					offset += 3
				}
			}
		} else if opcode == 7 {
			if modes[0] == 0 && modes[1] == 0 {
				if intCode[intCode[offset+1]] < intCode[intCode[offset+2]] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			} else if modes[0] == 0 && modes[1] == 1 {
				if intCode[intCode[offset+1]] < intCode[offset+2] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			} else if modes[0] == 1 && modes[1] == 0 {
				if intCode[offset+1] < intCode[intCode[offset+2]] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			} else if modes[0] == 1 && modes[1] == 1 {
				if intCode[offset+1] < intCode[offset+2] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			}
			offset += 4
		} else if opcode == 8 {
			if modes[0] == 0 && modes[1] == 0 {
				if intCode[intCode[offset+1]] == intCode[intCode[offset+2]] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			} else if modes[0] == 0 && modes[1] == 1 {
				if intCode[intCode[offset+1]] == intCode[offset+2] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			} else if modes[0] == 1 && modes[1] == 0 {
				if intCode[offset+1] == intCode[intCode[offset+2]] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			} else if modes[0] == 1 && modes[1] == 1 {
				if intCode[offset+1] == intCode[offset+2] {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 1
					} else {
						intCode[offset+3] = 1
					}
				} else {
					if modes[2] == 0 {
						intCode[intCode[offset+3]] = 0
					} else {
						intCode[offset+3] = 0
					}
				}
			}
			offset += 4
		} else {
			fmt.Println("error")
			os.Exit(1)
		}
	}
}

func main() {
	file, _ := os.Open("d05-input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intCodeStringArray := strings.Split(scanner.Text(), ",")
		intCode := stringArrayToIntArray(intCodeStringArray)
		intCodeComputer(intCode, 5)
	}
}
