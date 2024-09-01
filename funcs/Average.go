package Mathematics


func Average(Valeur []int) int {
	var somme int
	for i:=0;i<len(Valeur);i++ {
		somme+=Valeur[i]

	}
	return somme/len(Valeur)
}