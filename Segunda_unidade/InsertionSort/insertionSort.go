package main

func insertionSort(v []int) []int {
	for insercao := 1; insercao < len(v); insercao++ {
		for i := insercao; i >= 1 && v[i] < v[i-1]; i-- {
			v[i], v[i-1] = v[i-1], v[i]
		}
	}
	return v
}

func main() {

	v := []int{2, 8, 6, 10, 4, 5, 3}
	v = insertionSort(v)
	for i := 0; i < len(v); i++ {
		print(v[i], " ")
	}

}
