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
				buffer.WriteString(string(s[pair_of_j]))
			}
		}
	}
	return buffer.String()
}

func tailContribution(offset, level, tail int) int {
	ret := 0
	if tail > level {
		ret++
	}
	if level > 0 && level < offset-1 {
		if 2*offset-1-level >= tail {
			ret++
		}
	}
	return ret
}

func Decode(offset int, s string) string {
	jump := 2 * (offset - 1)
	num_vs := int(len(s) / jump)
	tail := len(s) - jump*num_vs
	level_ctr := make([]int, offset)
	for level := 0; level < offset; level++ {
		if level_ctr[level] = num_vs; level > 0 && level < offset-1 {
			level_ctr[level] *= 2
		}
		// Deal with the tail.
		level_ctr[level] += tailContribution(offset, level, tail)
	}
	deciphered := make([]byte, len(s))
	ctr := 0
	for level := 0; level < offset; level++ {
		for let := 0; let < level_ctr[level]; let++ {
			pos := level + jump*let
			if pos < len(s) {
				deciphered[pos] = s[ctr]
				ctr++
			}
			pair := pos + (jump - level*2)
			if pair > pos && pair < pos+jump && pair < len(s) {
				deciphered[pair] = s[ctr]
				ctr++
			}
		}
	}
	return string(deciphered)
}
