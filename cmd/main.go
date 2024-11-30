package main

import (
	"fmt"
	"math/rand"
)

const (
	FSHIFT  = 11
	FIXED_1 = 1 << FSHIFT
)

func main() {
	var (
		load   uint64 = 0
		exp    uint64 = 1884
		active uint64 = 0
	)

	var randValue int
	for i := 0; i < 220; i++ {
		if active != 21 { // When it reaches 21 let it be always 21
			randValue = rand.Intn(10)
			if randValue >= 4 { // 60% that active tasks amount will increase
				active++
			} else {
				if active != 0 {
					active--
				}
			}
		}

		result := calcLoad(load, exp, active*FIXED_1)
		load = result

		output := getLoadAverage(result, FIXED_1/200, 0)
		if i%2 == 0 {
			fmt.Printf("Load Average in 1min: %d.%d\tActive Tasks (R + D): %d\n", loadInt(output), loadFraq(output), active)
		}

	}
}

func calcLoad(load, exp uint64, active uint64) uint64 {
	newLoad := load*exp + active*(FIXED_1-exp)

	if active >= load {
		newLoad += FIXED_1 - 1
	}
	return newLoad / FIXED_1
}

func getLoadAverage(load, offset uint64, shift uint64) uint64 {
	return (load + offset) << shift
}

func loadInt(x uint64) uint64 {
	return x >> FSHIFT
}

func loadFraq(x uint64) uint64 {
	return loadInt(((x) & (FIXED_1 - 1)) * 100)
}
