package railfence

import "bytes"

func max(l, r int) int {
	if r > l {
		return r
	}
	return l
}

func Encode(offset int, s string) string {
	var buffer bytes.Buffer
	jump := 2 * (offset - 1)
	for level := 0; level < offset; level++ {
		for j := level; j < len(s); j += jump {
			buffer.WriteString(string(s[j]))
			pair_of_j := j + (jump - 2*level)
			if pair_of_j < len(s) && pair_of_j != j && pair_of_j < j+jump {
				buffer.WriteString(string(s[j]))
			}
		}
	}
	return buffer.String()
}
