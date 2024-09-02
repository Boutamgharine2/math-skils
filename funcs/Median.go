package Mathematics

import (
	"math"
	"sort"
)

func Median(Valeur []int) int {
	Valeur = Clasement(Valeur)
	sort.Ints(Valeur)
	n := len(Valeur)
	if n%2 == 0 {
		return int(math.Round(float64(Valeur[(n/2)-1]+Valeur[n/2]) / 2))
	}
	return Valeur[n/2]
}

func Clasement(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
