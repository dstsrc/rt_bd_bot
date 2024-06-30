package run

import "testing"

type rect struct {
	a int
	b int
}

func TestFind(t *testing.T) {

	r := rect{4, 5}

	want := 9

	got := RunS(r)

	if want != got {
		t.Errorf("Find() got = %v, want %v", got, want)
	}
}

func (r rect) FindS() int {
	return r.a + r.b
}
