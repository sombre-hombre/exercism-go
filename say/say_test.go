package say

import (
	"strings"
	"testing"
)

func TestSay(t *testing.T) {
	for _, tc := range testCases {
		actual, ok := Say(tc.input)
		if tc.expectError {
			if ok {
				t.Fatalf("FAIL: %s\nExpected error but received: %v", tc.description, actual)
			}
		} else if !ok {
			t.Fatalf("FAIL: %s\nDid not expect an error", tc.description)
		} else if actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %v\nActual: %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestWriteTens(t *testing.T) {
	for _, tc := range tensCases {
		var sb strings.Builder
		writeTens(tc.input, &sb)
		actual := sb.String()
		if actual != tc.expected {
			t.Fatalf("Expected: %v\nActual: %v", tc.expected, actual)
		}
	}
}

func TestWriteHundreds(t *testing.T) {
	for _, tc := range hundredsCases {
		var sb strings.Builder
		writeHundreds(tc.input, &sb)
		actual := sb.String()
		if actual != tc.expected {
			t.Fatalf("Expected: %v\nActual: %v", tc.expected, actual)
		}
	}
}

func BenchmarkSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Say(tc.input)
		}
	}
}
