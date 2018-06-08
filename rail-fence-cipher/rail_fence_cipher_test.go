package railfence

import (
	"reflect"
	"testing"
)

func testCases(op func(string, int) string, cases []testCase, t *testing.T) {
	for _, tc := range cases {
		if actual := op(tc.message, tc.rails); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %q\nActual: %q", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestEncode(t *testing.T) { testCases(Encode, encodeTests, t) }
func TestDecode(t *testing.T) { testCases(Decode, decodeTests, t) }

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range encodeTests {
			Encode(tc.message, tc.rails)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range decodeTests {
			Decode(tc.message, tc.rails)
		}
	}
}

func TestGetKey(t *testing.T) {
	expected := []int{0, 10, 1, 9, 11, 2, 8, 12, 3, 7, 13, 4, 6, 14, 16, 5, 15}

	mx := getKey(6, 17)

	if !reflect.DeepEqual(expected, mx) {
		t.Fatalf("makeMatrix(%d, %d) = %v, want %v",
			4, 7, mx, expected)
	}
}
