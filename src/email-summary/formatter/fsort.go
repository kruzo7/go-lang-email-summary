package formatter

import (
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type ByValue []Pair

func (a ByValue) Len() int {
	return len(a)
}

func (a ByValue) Less(i, j int) bool {
	return a[i].Value > a[j].Value
}

func (a ByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (f *Formatter) SortEmailsByValueDesc() []Pair {

	pairs := make([]Pair, 0, len(f.emailsnumber))

	for k, v := range f.emailsnumber {
		pairs = append(pairs, Pair{Key: k, Value: v})
	}

	sort.Sort(ByValue(pairs))

	return pairs
}
