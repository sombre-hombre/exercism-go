package pythagorean

// Use this type definition,
//
//    type Triplet [3]int
//
// and implement two functions,
//
//    Range(min, max int) []Triplet
//    Sum(p int) []Triplet
//
// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
//
// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
//
// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.

import (
	"reflect"
	"testing"
)

var rangeTests = []struct {
	min, max int
	ts       []Triplet
}{
	{1, 10, []Triplet{{3, 4, 5}, {6, 8, 10}}},
	{11, 20, []Triplet{{12, 16, 20}}},
}

func TestRange(t *testing.T) {
	for _, test := range rangeTests {
		ts := Range(test.min, test.max)
		if !reflect.DeepEqual(ts, test.ts) {
			t.Fatalf("Range(%d, %d) = %v, want %v",
				test.min, test.max, ts, test.ts)
		}
	}
}

var coprimeTests = []struct {
	x, y int
	gcd  int
}{
	{8, 15, 1},
	{49, 15, 1},
	{1071, 462, 21},
	{6, 9, 3},
}

func TestGCD(t *testing.T) {
	for _, test := range coprimeTests {
		actual := getGDC(test.x, test.y)
		if actual != test.gcd {
			t.Fatalf("getGDC(%d, %d) = %d, want %d",
				test.x, test.y, actual, test.gcd)
		}
	}
}

var primitiveTripletsTests = []struct {
	min, max int
	ts       []Triplet
}{
	{0, 25, []Triplet{{3, 4, 5}, {8, 15, 17}, {5, 12, 13}, {7, 24, 25}}},
	{0, 53, []Triplet{
		{3, 4, 5}, {8, 15, 17}, {12, 35, 37}, {5, 12, 13},
		{20, 21, 29}, {28, 45, 53}, {7, 24, 25}, {9, 40, 41},
	}},
	{100, 200, []Triplet{{104, 153, 185}, {119, 120, 169}}},
}

func TestGetPrimitiveTriplet(t *testing.T) {
	for _, test := range primitiveTripletsTests {
		ts := getPrimitiveTriplet(test.min, test.max)
		if !reflect.DeepEqual(ts, test.ts) {
			t.Fatalf("getPrimitiveTriplet(%d, %d) = %v, want %v",
				test.min, test.max, ts, test.ts)
		}
	}
}

var sumTests = []struct {
	sum int
	ts  []Triplet
}{
	{180, []Triplet{{18, 80, 82}, {30, 72, 78}, {45, 60, 75}}},
	{1000, []Triplet{{200, 375, 425}}},
}

func TestSum(t *testing.T) {
	for _, test := range sumTests {
		ts := Sum(test.sum)
		if !reflect.DeepEqual(ts, test.ts) {
			t.Fatalf("Sum(%d) = %v, want %v",
				test.sum, ts, test.ts)
		}
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Range(1, 100)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1000)
	}
}
