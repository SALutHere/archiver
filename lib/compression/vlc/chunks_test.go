package vlc

import (
	"reflect"
	"testing"
)

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

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		input    BinaryChunks
		expected string
	}{
		{BinaryChunks{}, ""},
		{BinaryChunks{"10101010"}, "10101010"},
		{BinaryChunks{"10101010", "11110000"}, "1010101011110000"},
		{BinaryChunks{"00000000", "00000001", "00000010"}, "000000000000000100000010"},
		{BinaryChunks{""}, ""},
		{BinaryChunks{"1", "0", "1"}, "101"},
	}

	for _, tt := range tests {
		result := tt.input.Join()
		if result != tt.expected {
			t.Errorf("BinaryChunks(%v).Join() = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "base test",
			data: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
