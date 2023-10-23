package main

import (
	"fmt"
)

type Process struct {
	Name  string
	Burst int
}

func roundRobin(processes []Process, timeSlice int) {
	queue := make([]Process, len(processes))
	copy(queue, processes)

	currentTime := 0

	for len(queue) > 0 {
		currentProcess := queue[0]
		queue = queue[1:]

		if currentProcess.Burst > timeSlice {
			fmt.Printf("Executing %s for %d units of time\n", currentProcess.Name, timeSlice)
			currentProcess.Burst -= timeSlice
			currentTime += timeSlice
			queue = append(queue, currentProcess)
		} else {
			fmt.Printf("%s has completed its execution at time %d\n", currentProcess.Name, currentTime+currentProcess.Burst)
			currentTime += currentProcess.Burst
		}
	}
}

func main() {

	processes := []Process{
		{"P1", 24},
		{"P2", 11},
		{"P3", 3},
	}

	timeSlice := 4

	roundRobin(processes, timeSlice)
}
