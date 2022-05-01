package sol

func minEatingSpeed(piles []int, h int) int {
	// find max
	max := 0
	for _, val := range piles {
		if max < val {
			max = val
		}
	}
	min := 0
	for min < max {
		mid := (min + max) / 2
		hSpent := 0
		for _, val := range piles {
			hSpent += val / mid
			if val%mid != 0 {
				hSpent += 1
			}
		}
		if hSpent <= h {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return max
}
