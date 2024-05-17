package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKmp(t *testing.T) {
	type args struct {
		str  string
		word string
	}

	tests := []struct {
		name           string
		args           args
		expectedResult []int
	}{
		{
			name: "The 1st standart test",
			args: args{
				str:  "ABC ABCDAB ABCDABCDABDE",
				word: "ABCDABD",
			},
			expectedResult: []int{15},
		},
		{
			name: "The 2nd standart test",
			args: args{
				str:  "this is simple test text... other test text",
				word: "text",
			},
			expectedResult: []int{20, 39},
		},
		{
			name: "The 3rd standart test",
			args: args{
				str:  "Aaaaaaaaaaaaaab aaaaa",
				word: "Aaaaaaaaaaaaaaa",
			},
			expectedResult: []int{},
		},
		{
			name: "Negative test",
			args: args{
				str:  "",
				word: "test",
			},
			expectedResult: []int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := Kmp(test.args.str, test.args.word)
			assert.EqualValues(tt, test.expectedResult, result)
		})
	}
}

func TestKmpTable(t *testing.T) {
	type args struct {
		word []rune
	}

	tests := []struct {
		name           string
		args           args
		expectedResult []int
	}{
		{
			name: "One match",
			args: args{
				word: []rune("abfgabh"),
			},
			//           +----+--------+--------+--------+-------+-------+--------+
			// word      | a  |   b    |   f    |   g    |   a   |   b   |   h    |
			//           +----+--------+--------+--------+-------+---+---+--------+---+
			// table     | -1 |   0    |   0    |   0    |  -1   |   0   |   2    | 0 |
			//           +----+--------+--------+--------+-------+-------+--------+---+
			// position  |    |   1    |   2    |   3    |   4   |   5   |   6    |
			//           +----+--------+--------+--------+-------+-------+--------+
			// candidate |    |   0    |   0    |   0    | (0)++ | (1)++ |   0    |
			//           |    | (-1)++ | (-1)++ | (-1)++ |       |       | (-1)++ |
			//           +----+--------+--------+--------+-------+-------+--------+
			expectedResult: []int{-1, 0, 0, 0, -1, 0, 2, 0},
		},
		{
			name: "Two match",
			args: args{
				word: []rune("abcdabc abc"),
			},
			//           +----+--------+--------+--------+-------+-------+-------+--------+-------+-------+-------+
			// word      | a  |   b    |   c    |   d    |   a   |   b   |   c   |   a    |   a   |   b   |   c   |
			//           +----+--------+--------+--------+-------+---+---+-------+--------+-------+-------+-------+---+
			// table     | -1 |   0    |   0    |   0    |  -1   |   0   |   0   |   3    |  -1   |   0   |   0   | 3 |
			//           +----+--------+--------+--------+-------+-------+-------+--------+-------+-------+-------+---+
			// position  |    |   1    |   2    |   3    |   4   |   5   |   6   |   7    |   8   |   9   |  10   |
			//           +----+--------+--------+--------+-------+-------+-------+--------+-------+-------+-------+
			// candidate |    |   0    |   0    |   0    | (0)++ | (1)++ | (2)++ |   0    |   0   |  1++  |  2++  |
			//           |    | (-1)++ | (-1)++ | (-1)++ |       |       |       | (-1)++ |       |       |       |
			//           +----+--------+--------+--------+-------+-------+-------+--------+-------+-------+-------+
			expectedResult: []int{-1, 0, 0, 0, -1, 0, 0, 3, -1, 0, 0, 3},
		},
		{
			name: "Two match",
			args: args{
				word: []rune("abcdabcdabcd"),
			},
			//           +----+--------+--------+--------+-------+-------+-------+-------+-------+-------+-------+-------+
			// word      | a  |   b    |   c    |   d    |   a   |   b   |   c   |   d   |   a   |   b   |   c   |   d   |
			//           +----+--------+--------+--------+-------+---+---+-------+-------+-------+-------+-------+-------+---+
			// table     | -1 |   0    |   0    |   0    |  -1   |   0   |   0   |   0   |  -1   |   0   |   0   |   0   | 8 |
			//           +----+--------+--------+--------+-------+-------+-------+-------+-------+-------+-------+-------+---+
			// position  |    |   1    |   2    |   3    |   4   |   5   |   6   |   7   |   8   |   9   |  10   |  11   |
			//           +----+--------+--------+--------+-------+-------+-------+-------+-------+-------+-------+-------+
			// candidate |    |   0    |   0    |   0    | (0)++ | (1)++ | (2)++ | (3)++ | (4)++ | (5)++ | (6)++ | (7)++ |
			//           |    | (-1)++ | (-1)++ | (-1)++ |       |       |       |       |       |       |       |       |
			//           +----+--------+--------+--------+-------+-------+-------+-------+-------+-------+-------+-------+
			expectedResult: []int{-1, 0, 0, 0, -1, 0, 0, 0, -1, 0, 0, 0, 8},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := kmpTable(test.args.word)
			assert.EqualValues(tt, test.expectedResult, result)
		})
	}
}
