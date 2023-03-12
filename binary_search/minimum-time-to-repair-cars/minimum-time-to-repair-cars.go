package minimum_time_to_repair_cars

import (
	"algorithm/utils"
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	// sum(sqrt(minutes / r)) >= numbers of car
	check := func(i int64) bool {
		sum := 0
		for _, v := range ranks {
			sum += int(math.Sqrt(float64(i) / float64(v)))
		}
		return sum >= cars
	}

	minMinutes := int64(utils.Min(ranks...))
	maxMinutes := int64(utils.Min(ranks...) * cars * cars)
	for minMinutes <= maxMinutes {
		midMinutes := (maxMinutes-minMinutes)/2 + minMinutes
		if check(midMinutes) {
			maxMinutes = midMinutes - 1
		} else {
			minMinutes = midMinutes + 1
		}
	}
	return minMinutes
}

// https://leetcode.com/problems/minimum-time-to-repair-cars/description/
// 二分法
