package treesort_test

import (
	"math/rand"
	"sort"
	"testing"
	"github.com/the_go_programming_language/chapter04/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
