package main

import "testing"

type Test struct {
	in  []int
	out bool
}

var testsCases = []Test{
	{
		// 1
		[]int{1, 2, 9, 9, 29, 9, 3, 4},
		false,
	},
	{
		[]int{1, 3, 2},
		true,
	},
	{
		[]int{1, 3, 2},
		true,
	},
	{
		[]int{1, 2, 9, 3, 4},
		true,
	},
	{
		[]int{1, 2, 1, 2},
		false,
	},
	{
		[]int{1, 999, 2, 7},
		true,
	},
	{
		[]int{1, 2, 3, 4, 3, 6},
		true,
	},
}

func TestAlmostIncreasingSequence(t *testing.T) {

	for i, test := range testsCases {
		result := almostIncreasingSequence(test.in)
		if result != test.out {
			t.Errorf("Error in the test N°%v, Expected %v, but got %v. For the input %v", i+1, test.out, result, test.in)
		}
	}

}
