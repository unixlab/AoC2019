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

func (om OrbitMap) getWay(start *Orbit, dest string) ([]string, int) {
	var way []string
	var steps int
	for start.Name != dest && start.InOrbit != nil {
		way = append(way, start.Name)
		steps++
		start = start.InOrbit
	}
	way = append(way, start.Name)
	steps--
	return way, steps
}

func findClosestMeetingPoint(way1 []string, way2 []string) string {
	for _, sw1 := range way1 {
		for _, sw2 := range way2 {
			if sw1 == sw2 {
				return sw1
			}
		}
	}
	return ""
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

	myWay, _ := orbitMap.getWay(orbitMap.getOrbitByName("YOU"), "COM")
	santasWay, _ := orbitMap.getWay(orbitMap.getOrbitByName("SAN"), "COM")
	meetingPoint := findClosestMeetingPoint(myWay, santasWay)

	_, myDistance := orbitMap.getWay(orbitMap.getOrbitByName("YOU"), meetingPoint)
	_, santasDistance := orbitMap.getWay(orbitMap.getOrbitByName("SAN"), meetingPoint)

	fmt.Println(myDistance + santasDistance)
}
