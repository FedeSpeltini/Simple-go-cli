package expenses

func Average(expens ...float32) float32 {

	return Sum(expens...) / float32(len(expens))
}

func Sum(expens ...float32) float32 {
	var sum float32
	for _, v := range expens {
		sum += v
	}
	return sum
}

func Max(expens ...float32) float32 {
	var max float32

	if len(expens) == 0 {
		return 0
	}

	for _, v := range expens {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(expens ...float32) float32 {
	min := expens[0]

	if len(expens) == 0 {
		return 0
	}

	for _, exp := range expens {
		if exp < min {
			min = exp
		}
	}
	return min
}
