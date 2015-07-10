package railfence

import (
	"bytes"
	"sort"
    "unicode/utf8"
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
	char rune
}

type ByIndex []Tuple

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].ind < a[j].ind }

func zipSortUnzip(s string, order []int) string {
    s_len := utf8.RuneCountInString(s)
	tuples := make([]Tuple, s_len)
    index := 0
	for _, el := range s {
		tuples[index].ind = order[index]
		tuples[index].char = el
        index++
	}
	sort.Stable(ByIndex(tuples))
	var buffer bytes.Buffer
	for _, e := range tuples {
		buffer.WriteRune(e.char)
	}
	return buffer.String()
}

func Encode(offset int, s string) string {
	pattern := rail(offset, utf8.RuneCountInString(s))
	return zipSortUnzip(s, pattern)
}

func Decode(offset int, s string) string {
    s_len := utf8.RuneCountInString(s)
	var rail_str string
	for i := 0; i < s_len; i++ {
		rail_str += string(i)
	}
	enc_rail_str := Encode(offset, rail_str)
	enc_rail := make([]int, s_len)
	for i := 0; i < s_len; i++ {
		enc_rail[i] = int(enc_rail_str[i])
	}
	return zipSortUnzip(s, enc_rail)
}
