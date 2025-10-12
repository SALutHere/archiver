package vlc

import (
	"testing"
)

func TestPrepareText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"My name is Sergey", "!my name is !sergey"},
		{"HELLO", "!h!e!l!l!o"},
		{"hello", "hello"},
		{"HeLLo WoRLd", "!he!l!lo !wo!r!ld"},
		{"", ""},
		{"123", "123"},
		{"A b C", "!a b !c"},
		{"!Already!", "!!already!"},
	}

	for _, tt := range tests {
		result := prepareText(tt.input)
		if result != tt.expected {
			t.Errorf("prepareText(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
func TestEncodeBin(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "011"},
		{"!a", "001000011"},
		{"hello", "001110100100100100110001"},
		{"!he!l!lo", "001000001110100100000100100100000100110001"},
		{" ", "11"},
		{"t n", "10011110000"},
		{"", ""},
		{"xyz", "000000000010000001000000000000"},
		{"!b", "0010000000010"},
	}

	for _, tt := range tests {
		result := encodeBin(tt.input)
		if result != tt.expected {
			t.Errorf("encodeBin(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
func TestSplitByChunks(t *testing.T) {
	tests := []struct {
		bStr      string
		chunkSize int
		expected  BinaryChunks
	}{
		{"10101010", 8, BinaryChunks{"10101010"}},
		{"1010101010101010", 8, BinaryChunks{"10101010", "10101010"}},
		{"101010101010101", 8, BinaryChunks{"10101010", "10101010"}},
		{"101", 8, BinaryChunks{"10100000"}},
		{"", 8, BinaryChunks{}},
		{"111100001111", 4, BinaryChunks{"1111", "0000", "1111"}},
		{"1", 1, BinaryChunks{"1"}},
		{"101010", 3, BinaryChunks{"101", "010"}},
		{"1010101", 4, BinaryChunks{"1010", "1010"}},
	}

	for _, tt := range tests {
		result := splitByChunks(tt.bStr, tt.chunkSize)
		if len(result) != len(tt.expected) {
			t.Errorf("splitByChunks(%q, %d) length = %d; want %d", tt.bStr, tt.chunkSize, len(result), len(tt.expected))
			continue
		}
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("splitByChunks(%q, %d)[%d] = %q; want %q", tt.bStr, tt.chunkSize, i, result[i], tt.expected[i])
			}
		}
	}
}
func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		input    BinaryChunks
		expected HexChunks
	}{
		{BinaryChunks{"00000000"}, HexChunks{"00"}},
		{BinaryChunks{"00000001"}, HexChunks{"01"}},
		{BinaryChunks{"11111111"}, HexChunks{"FF"}},
		{BinaryChunks{"10101010"}, HexChunks{"AA"}},
		{BinaryChunks{"00001111", "11110000"}, HexChunks{"0F", "F0"}},
		{BinaryChunks{"00000010", "00000011"}, HexChunks{"02", "03"}},
		{BinaryChunks{}, HexChunks{}},
	}

	for _, tt := range tests {
		result := tt.input.ToHex()
		if len(result) != len(tt.expected) {
			t.Errorf("BinaryChunks(%v).ToHex() length = %d; want %d", tt.input, len(result), len(tt.expected))
			continue
		}
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("BinaryChunks(%v).ToHex()[%d] = %q; want %q", tt.input, i, result[i], tt.expected[i])
			}
		}
	}
}
func TestEncode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "3A 49 31"},
		{"HELLO", "20 C8 A4 12 41 24 44"},
		{"HeLLo WoRLd", "20 E9 04 90 4C 72 01 C4 84 10 49 40"},
		{"", ""},
		{"a", "60"},
		{"A", "21 80"},
		{"My name is Sergey", "20 30 3C 18 77 4A E4 2D 40 4A 04"},
		{"!Already!", "20 86 4A 2B 28 12 00"},
	}

	for _, tt := range tests {
		result := Encode(tt.input)
		if result != tt.expected {
			t.Errorf("Encode(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
