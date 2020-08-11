package uid

import "testing"

func TestUid(t *testing.T)  {
	// case 1: length > 0
	lists1 := []int{4, 7, 9, 10, 16, 32}
	for _, v := range lists1{
		if len(Uid(v)) != v {
			t.Errorf("expected length equal " + string(v))
		}
	}
	// case 2: length <= 0
	lists2 := []int{0, -10, -50, -7}
	defaultLength := 16
	for _, v := range lists2{
		if len(Uid(v)) != defaultLength {
			t.Errorf("expected length equal " + string(defaultLength))
		}
	}
}
