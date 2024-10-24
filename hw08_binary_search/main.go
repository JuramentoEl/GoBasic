package main

func main() {
}

func BinarySearch(s []int, i int) int {
	var findIndex int

	end := len(s)
	start := 0
	if end == 0 || s[end-1] < i || s[0] > i {
		return -1
	}
	findIndex = (end - start) / 2
	for s[findIndex] != i && start != end {
		if s[findIndex] < i {
			start = findIndex + 1
		} else {
			end = findIndex
		}

		findIndex = start + ((end - start) / 2)
	}

	if s[findIndex] == i {
		return findIndex
	}
	return -1
}
