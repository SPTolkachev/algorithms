package main

import "fmt"

func main() {
	str := "ABC ABCDAB ABCDABCDABDE"
	word := "ABCDABD"
	positions := Kmp(str, word)

	fmt.Println("str: ", str)
	fmt.Println("word: ", word)
	fmt.Println("positions: ", positions)
}

// Kmp is realisation of Knuth–Morris–Pratt algorithm.
//
// Arguments:
//   - str -- the text to be searched;
//   - word -- the word sought.
//
// Results:
//   - number of positions.
func Kmp(str string, word string) []int {
	positions := []int{}
	strChars := []rune(str)
	strLen := len(strChars)
	wordChars := []rune(word)
	wordLen := len(wordChars)

	if strLen == 0 {
		return positions
	}

	sp := 0 // the position of the current character in argument `str`
	wp := 0 // the position of the current character in argument `word`
	table := kmpTable(wordChars)

	for sp < strLen {
		if wordChars[wp] == strChars[sp] {
			sp++
			wp++

			if wp == wordLen {
				positions = append(positions, sp-wp)
				wp = table[wp] // table[len(wordLen)] can't be -1
			}
		} else {
			wp = table[wp]
			if wp < 0 {
				sp++
				wp++
			}
		}
	}

	return positions
}

// kmpTable is realisation kmp table.
//
// Arguments:
//   - word -- the word to be analyzed.
//
// Results:
//   - the table to be filled.
func kmpTable(word []rune) []int {
	lenWord := len(word)
	if lenWord == 0 {
		return nil
	}

	table := make([]int, lenWord+1)
	position := 1  // the current position we are computing in table
	candidate := 0 // the zero-based index in `word` of the next character of the current candidate substring

	table[0] = -1

	for position < lenWord {
		if word[position] == word[candidate] {
			table[position] = table[candidate]
		} else {
			table[position] = candidate

			for candidate >= 0 && word[position] != word[candidate] {
				candidate = table[candidate]
			}
		}

		position++
		candidate++
	}

	table[position] = candidate // only needed when all word occurrences are searched

	return table
}
