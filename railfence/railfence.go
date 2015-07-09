package railfence

import (
	"bytes"
	"sort"
)

func rail(rail_size, length int) []int {
	pattern := make([]int, length)
	dir := 1
	pattern[0] = 0
	for i := 1; i < length; i++ {
		pattern[i] = pattern[i-1] + dir
		if pattern[i] == 0 || pattern[i] == rail_size-1 {
			dir = dir * -1
		}
	}
	return pattern
}

type Tuple struct {
	ind  int
	char string
}

type ByIndex []Tuple

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].ind < a[j].ind }

func zipSortUnzip(s string, order []int) string {
	tuples := make([]Tuple, len(s))
	for i, el := range order {
		tuples[i].ind = el
		tuples[i].char = string(s[i])
	}
	sort.Stable(ByIndex(tuples))
	var buffer bytes.Buffer
	for _, e := range tuples {
		buffer.WriteString(e.char)
	}
	return buffer.String()
}

func Encode(offset int, s string) string {
	pattern := rail(offset, len(s))
	return zipSortUnzip(s, pattern)
}

func Decode(offset int, s string) string {
	var rail_str string
	for i := 0; i < len(s); i++ {
		rail_str += string(i)
	}
	enc_rail_str := Encode(offset, rail_str)
	enc_rail := make([]int, len(s))
	for i := 0; i < len(enc_rail); i++ {
		enc_rail[i] = int(enc_rail_str[i])
	}
	return zipSortUnzip(s, enc_rail)
}
