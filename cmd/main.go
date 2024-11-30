package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		load   float64 = 0
		exp    float64 = 0.08
		active int64   = 0
	)

	var randValue int
	for i := 0; i < 1200; i++ {
		randValue = rand.Intn(10)
		if randValue >= 4 { // 60% that active tasks amount will increase
			active++
		} else {
			active--
			if active < 0 {
				active = 0
			}
		}

		result := calcLoad(load, exp, active)
		if i%2 == 0 {
			fmt.Printf("Load Average in 1min: %3.2f\tActive Tasks (R + D): %d\n", result, active)
		}

		load = result
	}
}

func calcLoad(load, exp float64, active int64) float64 {
	newLoad := load*exp + float64(active)*(1-exp)
	return newLoad
}
