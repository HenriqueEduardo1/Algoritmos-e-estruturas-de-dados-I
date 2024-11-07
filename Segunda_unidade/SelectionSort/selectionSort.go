package main

func selectionSort(v []int) []int {
	for varredura := 0; varredura < len(v)-1; varredura++ {
		iMenor := varredura
		for i := varredura + 1; i < len(v); i++ {
			if v[i] < v[iMenor] {
				iMenor = i
			}
		}
		v[iMenor], v[varredura] = v[varredura], v[iMenor]
	}
	return v
}

func main() {
	v := []int{2, 8, 6, 10, 4, 5, 3}

	v = selectionSort(v)

	for i := 0; i < len(v); i++ {
		print(v[i], " ")
	}
}
