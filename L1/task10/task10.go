package task10

import (
	"fmt"
	"math"
	"sort"
)

func groupTemperatures(temperatures []float64, interval float64) map[int][]float64 {
	groups := make(map[int][]float64)

	// Sort temperatures
	sort.Float64s(temperatures)

	for _, temp := range temperatures {
		var groupKey int
		// Check the interval for current group
		if temp < 0 {
			groupKey = int(math.Ceil(temp/interval) * interval) // round up
		} else {
			groupKey = int(math.Floor(temp/interval) * interval) // round down
		}

		groups[groupKey] = append(groups[groupKey], temp)
	}

	return groups
}

func Launch() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	interval := 10.0

	groups := groupTemperatures(temperatures, interval)

	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}
