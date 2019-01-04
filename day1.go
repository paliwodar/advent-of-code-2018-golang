package aoc2018

func Calibrate(changes []int) int {
	sum := 0
	for _,num := range changes {
		sum += num
	}
	return sum
}

func Calibrate2(changes []int) int {
	frequency := 0
	i := 0
	seen := make(map[int]bool)

	for !seen[frequency] {
		seen[frequency] = true
		frequency += changes[i]
		i = (i + 1) % len(changes)
	}
	return frequency
}
