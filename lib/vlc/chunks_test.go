package vlc

import "testing"

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
func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		input    string
		expected HexChunks
	}{
		{"", HexChunks{""}},
		{"AA", HexChunks{"AA"}},
		{"AA BB CC", HexChunks{"AA", "BB", "CC"}},
		{"01 02 03 04", HexChunks{"01", "02", "03", "04"}},
		{"FF", HexChunks{"FF"}},
		{"0A 0B", HexChunks{"0A", "0B"}},
		{"AA  BB  CC", HexChunks{"AA", "", "BB", "", "CC"}},
	}

	for _, tt := range tests {
		result := NewHexChunks(tt.input)
		if len(result) != len(tt.expected) {
			t.Errorf("NewHexChunks(%q) length = %d; want %d", tt.input, len(result), len(tt.expected))
			continue
		}
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("NewHexChunks(%q)[%d] = %q; want %q", tt.input, i, result[i], tt.expected[i])
			}
		}
	}
}
func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		input    HexChunk
		expected BinaryChunk
	}{
		{"00", "00000000"},
		{"01", "00000001"},
		{"FF", "11111111"},
		{"AA", "10101010"},
		{"0F", "00001111"},
		{"F0", "11110000"},
		{"02", "00000010"},
		{"03", "00000011"},
		{"7F", "01111111"},
		{"80", "10000000"},
		{"0A", "00001010"},
		{"B7", "10110111"},
	}

	for _, tt := range tests {
		result := tt.input.ToBinary()
		if result != tt.expected {
			t.Errorf("HexChunk(%q).ToBinary() = %q; want %q", tt.input, result, tt.expected)
		}
	}

	// Test invalid input (should panic)
	invalidInputs := []HexChunk{"ZZ", "G1", "123", ""}
	for _, input := range invalidInputs {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("HexChunk(%q).ToBinary() did not panic on invalid input", input)
				}
			}()
			_ = input.ToBinary()
		}()
	}
}
func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		input    HexChunks
		expected BinaryChunks
	}{
		{HexChunks{"00"}, BinaryChunks{"00000000"}},
		{HexChunks{"01"}, BinaryChunks{"00000001"}},
		{HexChunks{"FF"}, BinaryChunks{"11111111"}},
		{HexChunks{"AA"}, BinaryChunks{"10101010"}},
		{HexChunks{"0F", "F0"}, BinaryChunks{"00001111", "11110000"}},
		{HexChunks{"02", "03"}, BinaryChunks{"00000010", "00000011"}},
		{HexChunks{}, BinaryChunks{}},
	}

	for _, tt := range tests {
		result := tt.input.ToBinary()
		if len(result) != len(tt.expected) {
			t.Errorf("HexChunks(%v).ToBinary() length = %d; want %d", tt.input, len(result), len(tt.expected))
			continue
		}
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("HexChunks(%v).ToBinary()[%d] = %q; want %q", tt.input, i, result[i], tt.expected[i])
			}
		}
	}

	// Test invalid input (should panic)
	invalidInputs := []HexChunks{
		{"ZZ"},
		{"G1"},
		{"123"},
		{""},
		{"AA", "ZZ"},
	}
	for _, input := range invalidInputs {
		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("HexChunks(%v).ToBinary() did not panic on invalid input", input)
				}
			}()
			_ = input.ToBinary()
		}()
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
