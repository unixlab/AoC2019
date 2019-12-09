package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringArrayToIntArray(stringNumbers []string) []int64 {
	var intNumbers []int64
	for _, v := range stringNumbers {
		intNumber, _ := strconv.Atoi(v)
		intNumbers = append(intNumbers, int64(intNumber))
	}
	return intNumbers
}

func parseOptCode(instructionCode int64) (int64, []int64) {
	instructionString := strconv.FormatInt(instructionCode, 10)
	instructionLength := len(instructionString)
	opcode := 0
	if instructionLength > 1 {
		opcode, _ = strconv.Atoi(instructionString[instructionLength-2:])
	} else {
		return instructionCode, []int64{0, 0, 0}
	}
	instructionOffset := instructionLength - 2
	parameterModes := []string{"0", "0", "0"}
	for i := instructionOffset; i > 0; i-- {
		parameterModes[instructionOffset-i] = instructionString[i-1 : i]
	}
	return int64(opcode), stringArrayToIntArray(parameterModes)
}

func getValue(intCode []int64, pos int64, base int64, modes []int64) (int64, []int64) {
	mode, modes := modes[0], modes[1:]
	if mode == 0 {
		return intCode[intCode[pos]], modes
	} else if mode == 1 {
		return intCode[pos], modes
	} else if mode == 2 {
		return intCode[intCode[pos]+base], modes
	} else {
		fmt.Println("unknown mode")
		os.Exit(1)
	}
	return 0, []int64{0}
}

func setValue(intCode []int64, pos int64, base int64, modes []int64, val int64) []int64 {
	mode, modes := modes[0], modes[1:]
	if mode == 0 {
		intCode[intCode[pos]] = val
	} else if mode == 1 {
		intCode[pos] = val
	} else if mode == 2 {
		intCode[intCode[pos]+base] = val
	} else {
		fmt.Println("unknown mode")
		os.Exit(1)
	}
	return intCode
}

func intCodeComputer(intCode []int64, offset int64, base int64, input []int64) (int64, int64, int64, bool) {
	output := input[0]
	for {
		opcode, modes := parseOptCode(intCode[offset])
		if opcode == 99 {
			return output, offset, base, true
		} else if opcode == 1 {
			firstNumber, modes := getValue(intCode, offset+1, base, modes)
			secondNumber, modes := getValue(intCode, offset+2, base, modes)
			intCode = setValue(intCode, offset+3, base, modes, firstNumber+secondNumber)
			offset += 4
		} else if opcode == 2 {
			firstNumber, modes := getValue(intCode, offset+1, base, modes)
			secondNumber, modes := getValue(intCode, offset+2, base, modes)
			intCode = setValue(intCode, offset+3, base, modes, firstNumber*secondNumber)
			offset += 4
		} else if opcode == 3 {
			if len(input) < 1 {
				fmt.Println("too few inputs")
				os.Exit(1)
			}
			intCode = setValue(intCode, offset+1, base, modes, input[0])
			input = input[1:]
			offset += 2
		} else if opcode == 4 {
			output, _ = getValue(intCode, offset+1, base, modes)
			offset += 2
			return output, offset, base, false
		} else if opcode == 5 {
			firstParam, modes := getValue(intCode, offset+1, base, modes)
			if firstParam != 0 {
				offset, modes = getValue(intCode, offset+2, base, modes)
			} else {
				offset += 3
			}
		} else if opcode == 6 {
			firstParam, modes := getValue(intCode, offset+1, base, modes)
			if firstParam == 0 {
				offset, modes = getValue(intCode, offset+2, base, modes)
			} else {
				offset += 3
			}
		} else if opcode == 7 {
			firstParam, modes := getValue(intCode, offset+1, base, modes)
			secondParam, modes := getValue(intCode, offset+2, base, modes)
			if firstParam < secondParam {
				intCode = setValue(intCode, offset+3, base, modes, 1)
			} else {
				intCode = setValue(intCode, offset+3, base, modes, 0)
			}
			offset += 4
		} else if opcode == 8 {
			firstParam, modes := getValue(intCode, offset+1, base, modes)
			secondParam, modes := getValue(intCode, offset+2, base, modes)
			if firstParam == secondParam {
				intCode = setValue(intCode, offset+3, base, modes, 1)
			} else {
				intCode = setValue(intCode, offset+3, base, modes, 0)
			}
			offset += 4
		} else if opcode == 9 {
			var increment int64
			increment = 0
			increment, modes = getValue(intCode, offset+1, base, modes)
			base += increment
			offset += 2
		} else {
			fmt.Println("error")
			os.Exit(1)
		}
	}
}

func main() {
	file, _ := os.Open("d09-input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intCodeStringArray := strings.Split(scanner.Text(), ",")
		intCode := make([]int64, 10000)
		for k, v := range stringArrayToIntArray(intCodeStringArray) {
			intCode[k] = v
		}

		done := false
		var offset, result, base int64
		offset = 0
		result = 0
		base = 0
		for {
			result, offset, base, done = intCodeComputer(intCode, offset, base, []int64{1})
			if !done {
				fmt.Println(result)
			} else {
				break
			}
		}
	}
}
