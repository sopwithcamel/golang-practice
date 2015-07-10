package railfence

import "testing"

func TestEncode(t *testing.T) {
	cases := []struct {
		in     string
		offset int
		want   string
	}{
		{"REDDITCOMRDAILYPROGRAMMER", 3, "RIMIRAREDTORALPORMEDCDYGM"},
		{"LOLOLOLOLOLOLOLOLO", 2, "LLLLLLLLLOOOOOOOOO"},
		{"THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG", 4, "TCNMRZHIKWFUPETAYEUBOOJSVHLDGQRXOEO"},
		{"3141592653589793238462643383279502884197169399375105820974944592307816406286", 7, "3934546187438171450245968893099481332327954266552620198731963475632908289907"},
        {"日本語日本語", 3, "日本本日語語"},
	}
	for _, c := range cases {
		got := Encode(c.offset, c.in)
		if got != c.want {
			t.Errorf("Encode(%q, %q) == %q, want %q", c.offset, c.in, got, c.want)
		}
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		in     string
		offset int
		want   string
	}{
		{"RIMIRAREDTORALPORMEDCDYGM", 3, "REDDITCOMRDAILYPROGRAMMER"},
		{"LLLLLLLLLOOOOOOOOO", 2, "LOLOLOLOLOLOLOLOLO"},
		{"TCNMRZHIKWFUPETAYEUBOOJSVHLDGQRXOEO", 4, "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"},
		{"3934546187438171450245968893099481332327954266552620198731963475632908289907", 7, "3141592653589793238462643383279502884197169399375105820974944592307816406286"},
        {"日本本日語語", 3, "日本語日本語"},
	}
	for _, c := range cases {
		got := Decode(c.offset, c.in)
		if got != c.want {
			t.Errorf("Decode(%q, %q) == %q, want %q", c.offset, c.in, got, c.want)
		}
	}
}
