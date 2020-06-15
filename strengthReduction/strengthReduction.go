package strengthReduction

// base 10
func DigitCount(v uint64) uint32 {
	var result uint32

	for v > 0 {
		result++
		v /= 10
	}
	return result
}

func StrengthReductionDigitCount(v uint64) uint32 {
	var result uint32 = 1

	for  {
		if v < 10 {
			return result
		}
		if v < 100 {
			return result + 1
		}
		if v < 1000 {
			return result + 2
		}
		if v < 10000 {
			return result + 3
		}
		v /= 10000
		result += 4
	}
}
