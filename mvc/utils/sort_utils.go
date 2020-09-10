package utils

import "sort"

// []int {9,8,7,6,5,11}

func BubbleSort(elements []int) []int {
	isFinish := true
	for isFinish {
		isFinish = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				isFinish = true
			}
		}
	}

	return elements
}

func CustomSort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}
