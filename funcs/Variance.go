package Mathematics

func Variance(Valeurs []int, moyenne int) int {
	n := 0
	for i := 0; i < len(Valeurs); i++ {
		n += (Valeurs[i] - moyenne) * (Valeurs[i] - moyenne)
	}

	return n / len(Valeurs)
}
