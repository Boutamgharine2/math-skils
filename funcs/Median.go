package Mathematics

func Median(Valeur []int) int {
	ordonè := Clasement(Valeur)
	if len(ordonè)%2 == 0 {
		return (ordonè[(len(ordonè)/2)-1] + ordonè[(len(ordonè)/2)]) / 2
	}
	return ordonè[((len(ordonè)+1)/2)-1]
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
