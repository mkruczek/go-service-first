package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubleSort(t *testing.T) {
	//given
	given := []int{3, 5, 1, 4, 2}

	//when
	bubleSort(given)

	//then
	assert.EqualValues(t, []int{1, 2, 3, 4, 5}, given)
}

func TestBubleSortSameValue(t *testing.T) {
	//given
	given := []int{1, 1, 1, 1, 1}

	//when
	bubleSort(given)

	//then
	assert.EqualValues(t, []int{1, 1, 1, 1, 1}, given)
}

func TestBubleSortHappyScenario(t *testing.T) {
	//given
	given := []int{1, 2, 3, 4, 5}

	//when
bubleSort(given)

	//then
	assert.EqualValues(t, []int{1, 2, 3, 4, 5}, given)
}

func getElements(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[n-i-1] = i
	}
	return result
}

func BenchmarkSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		bubleSort(elements)
	}
}
func BenchmarkNativeGOSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkBubleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		bubleSort(elements)
	}
}
func BenchmarkNativeGOSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}
func BenchmarkBubleSort10000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		bubleSort(elements)
	}
}
func BenchmarkNativeGOSort10000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}
