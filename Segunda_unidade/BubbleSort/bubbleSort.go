package main

func bubbleSort(v []int) []int {
	for varredura := 1; varredura < len(v); varredura++ {
		trocou := false
		for i := 0; i < len(v)-varredura; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				trocou = true
			}
		}
		if !trocou {
			break
		}
	}
	return v
}

func main() {
	v := []int{2, 8, 6, 10, 4, 5, 3}
	v = bubbleSort(v)
	for i := 0; i < len(v); i++ {
		print(v[i], " ")
	}
}
