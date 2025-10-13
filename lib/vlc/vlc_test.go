package vlc

import (
	"reflect"
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
		name string
		str  string
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		encodedData []byte
		want        string
	}{
		{
			name:        "base test",
			encodedData: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.encodedData); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
