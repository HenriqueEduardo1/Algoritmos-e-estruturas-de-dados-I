package main

func merge(v []int, e []int, d []int) {
	indexV, indexE, indexD := 0, 0, 0
	for indexE < len(e) && indexD < len(d) {
		if e[indexE] < d[indexD] {
			v[indexV] = e[indexE]
			indexE++
		} else {
			v[indexV] = d[indexD]
			indexD++
		}
		indexV++
	}
	for indexE < len(e) {
		v[indexV] = e[indexE]
		indexE++
		indexV++
	}
	for indexD < len(d) {
		v[indexV] = d[indexD]
		indexD++
		indexV++
	}
}

func MergeSort(v []int) {

	if len(v) > 1 {
		meio := len(v) / 2
		e := make([]int, meio)
		d := make([]int, len(v)-meio)

		for i := 0; i < meio; i++ {
			e[i] = v[i]
		}
		for i := meio; i < len(v); i++ {
			d[i-meio] = v[i]
		}

		MergeSort(e)
		MergeSort(d)
		merge(v, e, d)
	}
}

func main() {
	v := []int{2, 8, 6, 10, 4, 5, 3}
	MergeSort(v)
	for i := 0; i < len(v); i++ {
		print(v[i], " ")
	}
}
