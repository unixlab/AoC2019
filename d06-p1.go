package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OrbitMap struct {
	Orbits []Orbit
}

func (om OrbitMap) getOrbitByName(name string) *Orbit {
	for _, v := range om.Orbits {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (om OrbitMap) getNumberOfOrbits(orbit *Orbit) int {
	numberOfOrbits := 0
	for orbit.InOrbit != nil {
		numberOfOrbits++
		orbit = orbit.InOrbit
	}
	return numberOfOrbits
}

type Orbit struct {
	Name    string
	InOrbit *Orbit
}

func main() {
	file, _ := os.Open("d06-input.txt")
	scanner := bufio.NewScanner(file)

	var orbitMap OrbitMap
	orbit := Orbit{"COM", nil}
	orbitMap.Orbits = append(orbitMap.Orbits, orbit)

	var orbitInput [][]string
	for scanner.Scan() {
		orbitInput = append(orbitInput, strings.Split(scanner.Text(), ")"))
	}

	for len(orbitInput) > 0 {
		var orbitRelation []string
		orbitRelation, orbitInput = orbitInput[0], orbitInput[1:]
		orbitAddr := orbitMap.getOrbitByName(orbitRelation[0])
		if orbitAddr == nil {
			orbitInput = append(orbitInput, orbitRelation)
		} else {
			orbit := Orbit{orbitRelation[1], orbitAddr}
			orbitMap.Orbits = append(orbitMap.Orbits, orbit)
		}
	}

	numberOfOrbits := 0
	for _, v := range orbitMap.Orbits {
		numberOfOrbits += orbitMap.getNumberOfOrbits(&v)
	}

	fmt.Println(numberOfOrbits)
}
