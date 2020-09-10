package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// Initialiation
	els := []int{9, 8, 7, 6, 5, 11}

	// Execution
	bubbleSortedEls := BubbleSort(els)

	// Assertion
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, len(els))
	assert.Equal(t, bubbleSortedEls, els)
}

func TestBubbleSortNilSlice(t *testing.T) {

	//sorted := BubbleSort(nil)

}

func getElement(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func TestGetElement(t *testing.T) {
	els := getElement(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 4, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 2, els[3])
	assert.EqualValues(t, 1, els[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElement(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort10(b *testing.B) {
	els := getElement(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElement(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElement(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElement(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElement(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
