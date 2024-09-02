package Mathematics

import "math"

func Variance(Valeurs []int, moyenne int) int {
	var n float64
	for i := 0; i < len(Valeurs); i++ {
		n += float64((Valeurs[i] - moyenne) * (Valeurs[i] - moyenne))
	}

	return int(math.Round(float64(n / float64(len(Valeurs)))))
}
