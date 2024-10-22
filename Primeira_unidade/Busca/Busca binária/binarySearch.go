package main

func BinarySearch(v []int, e int) int {
	ini := 0
	fim := len(v) - 1

	for ini <= fim {
		meio := (ini + fim) / 2

		if v[meio] == e {
			return meio
		} else if v[meio] < e {
			ini = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}

func BinarySearchRecursive(v []int, e int, ini int, fim int) int {
	if ini > fim {
		return -1
	}

	meio := (ini + fim) / 2

	if v[meio] == e {
		return meio
	} else {
		if e < v[meio] {
			return BinarySearchRecursive(v, e, ini, meio-1)
		} else {
			return BinarySearchRecursive(v, e, meio+1, fim)
		}
	}
}

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ini := 0
	fim := len(v) - 1

	println(BinarySearch(v, 1))
	println(BinarySearch(v, 2))
	println(BinarySearch(v, 5))
	println(BinarySearch(v, 9))
	println(BinarySearch(v, 10))
	println(BinarySearch(v, 20))
	println(BinarySearch(v, 0))

	println(BinarySearchRecursive(v, 1, ini, fim))
	println(BinarySearchRecursive(v, 2, ini, fim))
	println(BinarySearchRecursive(v, 5, ini, fim))
	println(BinarySearchRecursive(v, 9, ini, fim))
	println(BinarySearchRecursive(v, 10, ini, fim))
	println(BinarySearchRecursive(v, 20, ini, fim))
	println(BinarySearchRecursive(v, 0, ini, fim))
}
