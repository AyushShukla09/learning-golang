package main

import "testing"

func TestSum(t *testing.T) {

	type test struct {
		data []int
		ans int
	}

	tests := []test{
		test{[]int{2, -2}, 0},
		test{[]int{22, -2}, 20},
		test{[]int{20, -22}, -2},
		test{[]int{99999, 1}, 100000},
	}

	for _,v := range tests {
		x := sum(v.data...)
		if x != v.ans {
			t.Error("Expected ", v.ans, " Got ", x)
		}
	}

}

