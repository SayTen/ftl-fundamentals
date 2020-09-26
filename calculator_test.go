package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	name        string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Adds two positive numbers"},
		{a: -3, b: 2, want: -1, name: "Adds one negative and one positive number"},
		{a: -3, b: -3, want: -6, name: "Adds two negative numbers"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)

		if want != got {
			t.Fatalf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 0, name: "Subtracts two positive numbers"},
		{a: 4, b: -2, want: 6, name: "Subtracts negative number from positive number"},
		{a: -4, b: 2, want: -6, name: "Subtracts positive number from negative number"},
		{a: -4, b: -2, want: -2, name: "Subtracts negative number from negative number"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}

	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Multiplys two positive numbers"},
		{a: 1, b: 1, want: 1, name: "Multiplys number against one"},
		{a: 5, b: 0, want: 0, name: "Multiplys number against zero"},
		{a: 5, b: -5, want: -25, name: "Multiplys one positive and one negative number"},
		{a: -5, b: -5, want: 25, name: "Multiplys two negative numbers"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 6, b: 2, want: 3, name: "Divides two positive numbers"},
		{a: 6, b: -2, want: -3, name: "Divides a positive number by a negative number"},
		{a: -6, b: 2, want: -3, name: "Divides a negative number by a positive number"},
		{a: -6, b: -2, want: 3, name: "Divides two negative numbers"},
		{a: 6, b: 0, want: 999, name: "Divide by zero", errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s: Divide(%f, %f): unexpected error status: %v", tc.name, tc.a, tc.b, errReceived)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
