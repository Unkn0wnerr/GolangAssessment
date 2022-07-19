package math

import "testing"

type testAddArgs struct {
	num1     int
	num2     int
	expected int
}

type testSubtractArgs struct {
	num1     int
	num2     int
	expected int
}

var testCasesAdd = []testAddArgs{
	testAddArgs{1, 2, 3},
	testAddArgs{3, 4, 7},
	testAddArgs{5, 6, 11},
}

var testCasesSubtract = []testSubtractArgs{
	testSubtractArgs{2, 1, 1},
	testSubtractArgs{4, 3, 1},
	testSubtractArgs{6, 5, 1},
}

func TestAdd(t *testing.T) {
	for _, test := range testCasesAdd {
		got := Add(test.num1, test.num2)
		if got != test.expected {
			t.Errorf("Got %q, wanted %q", got, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, test := range testCasesSubtract {
		got := Subtract(test.num1, test.num2)
		if got != test.expected {
			t.Errorf("Got %q, wanted %q", got, test.expected)
		}
	}
}
