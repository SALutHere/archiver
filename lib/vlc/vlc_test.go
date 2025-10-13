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
	}

	for _, tt := range tests {
		result := Encode(tt.input)
		if result != tt.expected {
			t.Errorf("Encode(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
func TestDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3A 49 31", "hello"},
		{"20 C8 A4 12 41 24 44", "HELLO"},
		{"20 E9 04 90 4C 72 01 C4 84 10 49 40", "HeLLo WoRLd"},
		{"60", "a"},
		{"21 80", "A"},
		{"20 30 3C 18 77 4A E4 2D 40 4A 04", "My name is Sergey"},
	}

	for _, tt := range tests {
		result := Decode(tt.input)
		if result != tt.expected {
			t.Errorf("Decode(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
