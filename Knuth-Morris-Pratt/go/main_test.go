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
