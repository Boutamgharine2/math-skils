package Mathematics

import "math"

func Average(Valeur []int) int {
	var somme float64
	for i := 0; i < len(Valeur); i++ {
		somme += float64(Valeur[i])
	}
	return int(math.Round(somme / float64(len(Valeur))))
}
