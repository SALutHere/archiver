package vlc

import (
	"reflect"
	"testing"
)

func TestEncodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		input    encodingTable
		expected DecodingTree
	}{
		{
			input: encodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			expected: DecodingTree{
				Zero: &DecodingTree{
					One: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &DecodingTree{
					Zero: &DecodingTree{
						Zero: &DecodingTree{
							One: &DecodingTree{
								Value: "b",
							},
						},
					},
					One: &DecodingTree{
						Value: "a",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		result := tt.input.DecodingTree()
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("encodingTable.decodingTree() = %v, expected %v", result, tt.expected)
		}
	}
}
