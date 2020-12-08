package fibb

import (
	"testing"
)

func Testfibbo(t *testing.T) {
	Memory := make(map[int]int)
	n := Fibbo(8, &Memory)
	if n != 21 {
		t.Errorf("Failed")
	}
}
