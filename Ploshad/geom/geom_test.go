package geom

import (
	"testing"
)

func TestFindA(t *testing.T) {
	a := 5
	b := 6
	expect := 11

	got := FindA(a, b)

	if got != expect {
		t.Errorf("Find() got = %v, expect %v", got, expect)
	}

}
