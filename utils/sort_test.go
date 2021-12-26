package utils

import "testing"

func TestSortInts(t *testing.T) {
	ints := []interface{}{}
	ints = append(ints, 4)
	ints = append(ints, 1)
	ints = append(ints, 2)
	ints = append(ints, 3)

	Sort(ints, IntComparator)

	for i := 1; i < len(ints); i++ {
		if ints[i-1].(int) > ints[i].(int) {
			t.Errorf("Not sorted!")
		}
	}
}

func TestSortStrings(t *testing.T) {

	strings := []interface{}{}
	strings = append(strings, "d")
	strings = append(strings, "a")
	strings = append(strings, "b")
	strings = append(strings, "c")

	Sort(strings, StringComparator)

	for i := 1; i < len(strings); i++ {
		if strings[i-1].(string) > strings[i].(string) {
			t.Errorf("Not sorted!")
		}
	}
}
