package uniq

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqBasic(t *testing.T) {
	const caseCount = 8
	var optsTable = [caseCount]Options{
		Options{},
		Options{OutputFormat: Count},
		Options{OutputFormat: Repeated},
		Options{OutputFormat: Once},

		Options{IgnoreCase: true},
		Options{OutputFormat: Count, IgnoreCase: true},
		Options{OutputFormat: Repeated, IgnoreCase: true},
		Options{OutputFormat: Once, IgnoreCase: true},
	}

	var testTable = []struct {
		input    []string
		expected [caseCount][]string
	}{
		{
			[]string{"", "0", "00", "000", "0000", "00000", "000000", "0000000", "00000000", "000000000", "0000000000", "00000000000", "000000000000", "0000000000000", "00000000000000"},
			[caseCount][]string{
				[]string{"", "0", "00", "000", "0000", "00000", "000000", "0000000", "00000000", "000000000", "0000000000", "00000000000", "000000000000", "0000000000000", "00000000000000"},
				[]string{"1 ", "1 0", "1 00", "1 000", "1 0000", "1 00000", "1 000000", "1 0000000", "1 00000000", "1 000000000", "1 0000000000", "1 00000000000", "1 000000000000", "1 0000000000000", "1 00000000000000"},
				nil,
				[]string{"", "0", "00", "000", "0000", "00000", "000000", "0000000", "00000000", "000000000", "0000000000", "00000000000", "000000000000", "0000000000000", "00000000000000"},
				[]string{"", "0", "00", "000", "0000", "00000", "000000", "0000000", "00000000", "000000000", "0000000000", "00000000000", "000000000000", "0000000000000", "00000000000000"},
				[]string{"1 ", "1 0", "1 00", "1 000", "1 0000", "1 00000", "1 000000", "1 0000000", "1 00000000", "1 000000000", "1 0000000000", "1 00000000000", "1 000000000000", "1 0000000000000", "1 00000000000000"},
				nil,
				[]string{"", "0", "00", "000", "0000", "00000", "000000", "0000000", "00000000", "000000000", "0000000000", "00000000000", "000000000000", "0000000000000", "00000000000000"},
			},
		},
		{
			[]string{"A", "AA", "AA", "AAA", "AAAAa", "AAAAaa", "AAAAaa", "AAAaaa", "AAaaaa", "Aaaaaa", "aaAAaa", "aaAAaa", "aaaaaa", "aaaaaaa", "aaaaaaa", "aaaaaaa"},
			[caseCount][]string{
				[]string{"A", "AA", "AAA", "AAAAa", "AAAAaa", "AAAaaa", "AAaaaa", "Aaaaaa", "aaAAaa", "aaaaaa", "aaaaaaa"},
				[]string{"1 A", "2 AA", "1 AAA", "1 AAAAa", "2 AAAAaa", "1 AAAaaa", "1 AAaaaa", "1 Aaaaaa", "2 aaAAaa", "1 aaaaaa", "3 aaaaaaa"},
				[]string{"AA", "AAAAaa", "aaAAaa", "aaaaaaa"},
				[]string{"A", "AAA", "AAAAa", "AAAaaa", "AAaaaa", "Aaaaaa", "aaaaaa"},
				[]string{"A", "AA", "AAA", "AAAAa", "AAAAaa", "aaaaaaa"},
				[]string{"1 A", "2 AA", "1 AAA", "1 AAAAa", "8 AAAAaa", "3 aaaaaaa"},
				[]string{"AA", "AAAAaa", "aaaaaaa"},
				[]string{"A", "AAA", "AAAAa"},
			},
		},
		{
			[]string{"", "0", "00", "00", "1", "1", "1", "1", "111", "222", "222", "2222", "333", "4444", "44455", "555", "555", "6666", "6666", "67777", "67777", "9", "9", "9", "9", "9", "9", "99", "999", "999", "99999"},
			[caseCount][]string{
				[]string{"", "0", "00", "1", "111", "222", "2222", "333", "4444", "44455", "555", "6666", "67777", "9", "99", "999", "99999"},
				[]string{"1 ", "1 0", "2 00", "4 1", "1 111", "2 222", "1 2222", "1 333", "1 4444", "1 44455", "2 555", "2 6666", "2 67777", "6 9", "1 99", "2 999", "1 99999"},
				[]string{"00", "1", "222", "555", "6666", "67777", "9", "999"},
				[]string{"", "0", "111", "2222", "333", "4444", "44455", "99", "99999"},
				[]string{"", "0", "00", "1", "111", "222", "2222", "333", "4444", "44455", "555", "6666", "67777", "9", "99", "999", "99999"},
				[]string{"1 ", "1 0", "2 00", "4 1", "1 111", "2 222", "1 2222", "1 333", "1 4444", "1 44455", "2 555", "2 6666", "2 67777", "6 9", "1 99", "2 999", "1 99999"},
				[]string{"00", "1", "222", "555", "6666", "67777", "9", "999"},
				[]string{"", "0", "111", "2222", "333", "4444", "44455", "99", "99999"},
			},
		},
		{
			[]string{"A", "A", "AB", "AB", "ABB", "ABB", "ABb", "AbB", "Abb", "aBB", "aBb", "abB", "abb", "abb"},
			[caseCount][]string{
				[]string{"A", "AB", "ABB", "ABb", "AbB", "Abb", "aBB", "aBb", "abB", "abb"},
				[]string{"2 A", "2 AB", "2 ABB", "1 ABb", "1 AbB", "1 Abb", "1 aBB", "1 aBb", "1 abB", "2 abb"},
				[]string{"A", "AB", "ABB", "abb"},
				[]string{"ABb", "AbB", "Abb", "aBB", "aBb", "abB"},
				[]string{"A", "AB", "abb"},
				[]string{"2 A", "2 AB", "10 abb"},
				[]string{"A", "AB", "abb"},
				nil,
			},
		},
	}

	for _, testCase := range testTable {
		for index, opts := range optsTable {
			resultStrs, err := Uniq(testCase.input, opts)
			assert.Equal(t, testCase.expected[index], resultStrs, fmt.Sprintf("Incorrect result. Expect %v, got %v", testCase.expected[index], resultStrs))
			assert.Nil(t, err, "Error return value is not nil, but no errors expected. Error messege: %v", err)
		}
	}
}

func TestUniqSkipChars(t *testing.T) {
	const countCases = 10
	var testTable = []struct {
		input    []string
		expected [countCases][]string
	}{
		{
			[]string{"", "aaaaaaaaaaaa", "aaaaaaaaaaaa", "baaaaaaaaaaa", "bbaaaaaaaaaa", "bbbbaaaaaaaa", "bbbbaaaaaaaa", "bbbbbaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "cccccddddddd", "cccccddddddd", "cccccddddddd", "dddddddddddd", "dddddddddddd"},
			[countCases][]string{
				[]string{"", "aaaaaaaaaaaa", "baaaaaaaaaaa", "bbaaaaaaaaaa", "bbbbaaaaaaaa", "bbbbbaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "cccccddddddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbaaaaaaaaaa", "bbbbaaaaaaaa", "bbbbbaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "cccccddddddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbaaaaaaaa", "bbbbbaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "cccccddddddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbaaaaaaaa", "bbbbbaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "cccccddddddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbbaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "cccccddddddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbbbbbaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "ccccccccdddd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "dddddddddddd"},
				[]string{"", "aaaaaaaaaaaa", "bbbbbbbbbbaa", "bbbbbbbbbbba", "bbbbbbbbbbbb", "cccccccccccd", "ccccccccccdd", "dddddddddddd"},
			},
		},
		{
			[]string{"      11", "1 111 11", "11111111", "21111111", "222 1111", "222  211", "22221111", "22222211", "32111 11", "3222 222", "32222222"},
			[countCases][]string{
				[]string{"      11", "1 111 11", "11111111", "21111111", "222 1111", "222  211", "22221111", "22222211", "32111 11", "3222 222", "32222222"},
				[]string{"      11", "1 111 11", "11111111", "222 1111", "222  211", "22221111", "22222211", "32111 11", "3222 222", "32222222"},
				[]string{"      11", "1 111 11", "11111111", "222 1111", "222  211", "22221111", "22222211", "32111 11", "3222 222", "32222222"},
				[]string{"      11", "1 111 11", "11111111", "222 1111", "222  211", "22221111", "22222211", "32111 11", "3222 222", "32222222"},
				[]string{"      11", "1 111 11", "11111111", "222  211", "22221111", "22222211", "32111 11", "3222 222", "32222222"},
				[]string{"      11", "11111111", "222  211", "22221111", "22222211", "32111 11", "32222222"},
				[]string{"      11", "32222222"},
				[]string{"      11", "32222222"},
				[]string{"32222222"},
				[]string{"32222222"},
			},
		},
	}

	for _, testCase := range testTable {
		opts := Options{}
		for numSkipChars := 0; numSkipChars < countCases; numSkipChars++ {
			opts.SkipChars = numSkipChars
			resultStrs, err := Uniq(testCase.input, opts)
			assert.Equal(t, testCase.expected[numSkipChars], resultStrs, fmt.Sprintf("Incorrect result. Expect %v, got %v", testCase.expected[numSkipChars], resultStrs))
			assert.Nil(t, err, "Error return value is not nil, but no errors expected. Error messege: %v", err)
		}
	}
}

func TestUniqSkipFields(t *testing.T) {
	const countCases = 5
	var testTable = []struct {
		input    []string
		expected [countCases][]string
	}{
		{[]string{"", "a a a aaa", "bbb bbb   a aaa", "bbbb bbb a aaa", "bbbb bbb a aaa", "bbbb bbb a aaa", "c bbb a aaa", "c ccc c", "c ccc cccccccccc aaa", "d ccc c", "d ddd c", "d ddd d"},
			[countCases][]string{
				[]string{"", "a a a aaa", "bbb bbb   a aaa", "bbbb bbb a aaa", "c bbb a aaa", "c ccc c", "c ccc cccccccccc aaa", "d ccc c", "d ddd c", "d ddd d"},
				[]string{"", "a a a aaa", "bbb bbb   a aaa", "bbbb bbb a aaa", "c ccc c", "c ccc cccccccccc aaa", "d ccc c", "d ddd c", "d ddd d"},
				[]string{"", "a a a aaa", "c ccc c", "c ccc cccccccccc aaa", "d ccc c", "d ddd d"},
				[]string{"", "a a a aaa", "c ccc c", "c ccc cccccccccc aaa", "d ddd d"},
				[]string{"d ddd d"},
			},
		},

		{[]string{"1 1 1 1 1 111", "1 1 1 1 1 111", "11111 1 1 111", "1111111 1 111", "111111111 111", "22222 1 1 1 1 1 111", "222222222 1 1 1 1 1 111", "222222222 2222 1 1 1 1 111", "222222222 2222 222 2 1 1 111", "3 2222 222 2 1 1 111"},
			[countCases][]string{
				[]string{"1 1 1 1 1 111", "11111 1 1 111", "1111111 1 111", "111111111 111", "22222 1 1 1 1 1 111", "222222222 1 1 1 1 1 111", "222222222 2222 1 1 1 1 111", "222222222 2222 222 2 1 1 111", "3 2222 222 2 1 1 111"},
				[]string{"1 1 1 1 1 111", "11111 1 1 111", "1111111 1 111", "111111111 111", "22222 1 1 1 1 1 111", "222222222 2222 1 1 1 1 111", "3 2222 222 2 1 1 111"},
				[]string{"1 1 1 1 1 111", "11111 1 1 111", "1111111 1 111", "111111111 111", "22222 1 1 1 1 1 111", "3 2222 222 2 1 1 111"},
				[]string{"1 1 1 1 1 111", "11111 1 1 111", "1111111 1 111", "22222 1 1 1 1 1 111", "3 2222 222 2 1 1 111"},
				[]string{"1 1 1 1 1 111", "11111 1 1 111", "3 2222 222 2 1 1 111"},
			},
		},
	}

	for _, testCase := range testTable {
		opts := Options{}
		for numSkipFields := 0; numSkipFields < countCases; numSkipFields++ {
			opts.SkipFields = numSkipFields
			resultStrs, err := Uniq(testCase.input, opts)
			assert.Equal(t, testCase.expected[numSkipFields], resultStrs, fmt.Sprintf("Incorrect result. Expect %v, got %v", testCase.expected[numSkipFields], resultStrs))
			assert.Nil(t, err, "Error return value is not nil, but no errors expected. Error messege: %v", err)
		}
	}
}

func TestUniqFormatedOutputSkipFieldsChars(t *testing.T) {
	var testTable = []struct {
		opts     Options
		input    []string
		expected []string
	}{
		{Options{OutputFormat: 'c', SkipFields: 1, SkipChars: 2}, []string{"1 111 1111", "1 221 1111", "2 331 1111", "4 4 1 1111"}, []string{"4 4 4 1 1111"}},
		{Options{OutputFormat: 'c', SkipFields: 2, SkipChars: 3}, []string{"0 1 2 3 4 5 6", "1111111 1 133 4 5 6", "22222222222 222222222222 333 4 5 6"}, []string{"3 22222222222 222222222222 333 4 5 6"}},
		{Options{OutputFormat: 'c', SkipFields: 5, SkipChars: 5}, []string{"1 111 11111 1 1 12341 1 1", "2222222222 3 4 5555 6 66661 1 1", "3 3 3 3 3 3 331 1 1", "4 33 3 3 3 331 1 1", ""}, []string{"3 1 111 11111 1 1 12341 1 1", "1 4 33 3 3 3 331 1 1", "1 "}},

		{Options{OutputFormat: 'd', SkipFields: 1, SkipChars: 2}, []string{"1 111 1111", "1 221 1111", "2 331 1111", "4 4 1 1111", "4 4 4 4444", "4 4 4 4444", "5 4 4 4444", "7 7 7777 7", "7 7 777777"}, []string{"1 111 1111", "4 4 4 4444"}},
		{Options{OutputFormat: 'd', SkipFields: 2, SkipChars: 3}, []string{"00000000 000 334", "00000000 0010 444", "0 1 2 3 4 5 6", "1111111 1 133 4 5 6", "22222222222 222222222222 333 4 5 6", "22222222222 222222222222 333 4 5 6", "22222222222 222222222222 333 4 5 6", "99999 9 9 9 9 9 9 9", "999 9 9 9 9 9999 99 9"}, []string{"00000000 000 334", "0 1 2 3 4 5 6"}},
		{Options{OutputFormat: 'd', SkipFields: 5, SkipChars: 5}, []string{"1 111 11111 1 1 12341 1 1", "21", "2222222222 3 4 5555 6 66661 1 1", "3 3 3 3 3 3 331 1 1", "4 33 3 3 3 331 1 1"}, []string{"2222222222 3 4 5555 6 66661 1 1"}},

		{Options{OutputFormat: 'u', SkipFields: 1, SkipChars: 2}, []string{"1 111 1111", "1 221 1111", "2 331 1111", "4 4 1 1111", "4 4 4 4444", "4 4 4 4444", "5 4 4 4444", "7 7 7777 7", "7 7 777777"}, []string{"7 7 7777 7", "7 7 777777"}},
		{Options{OutputFormat: 'u', SkipFields: 2, SkipChars: 3}, []string{"00000000 000 334", "00000000 0010 444", "0 1 2 3 4 5 6", "1111111 1 133 4 5 6", "22222222222 222222222222 333 4 5 6", "22222222222 222222222222 333 4 5 6", "22222222222 222222222222 333 4 5 6", "99999 9 9 9 9 9 9 9", "999 9 9 9 9 9999 99 9"}, []string{"99999 9 9 9 9 9 9 9", "999 9 9 9 9 9999 99 9"}},
		{Options{OutputFormat: 'u', SkipFields: 5, SkipChars: 5}, []string{"1 111 11111 1 1 12341 1 1", "21", "2222222222 3 4 5555 6 66661 1 1", "3 3 3 3 3 3 331 1 1", "4 33 3 3 3 331 1 1"}, []string{"1 111 11111 1 1 12341 1 1", "21", "4 33 3 3 3 331 1 1"}},
	}

	for _, testCase := range testTable {
		resultStrs, err := Uniq(testCase.input, testCase.opts)
		assert.Equal(t, testCase.expected, resultStrs, fmt.Sprintf("Incorrect result. Options: %v. Expect %v, got %v", testCase.opts, testCase.expected, resultStrs))
		assert.Nil(t, err, "Error return value is not nil, but no errors expected. Error messege: %v", err)
	}
}

func TestUniqInvalidArgs(t *testing.T) {
	var testTable = []struct {
		opts     Options
		input    []string
		expected []string
		errMsg   string
	}{
		{Options{OutputFormat: 'x'}, []string{"A", "A", "AB", "AB", "ABB", "ABB", "ABb", "AbB"}, nil, "no such output format"},
		{Options{}, nil, nil, "nil passed as slice"},
		{Options{SkipChars: -1}, []string{"A", "A", "AB", "AB", "ABB", "ABB", "ABb", "AbB"}, nil, "passed negative number of characters to skip"},
		{Options{SkipFields: -1}, []string{"A", "A", "AB", "AB", "ABB", "ABB", "ABb", "AbB"}, nil, "passed negative number of fields to skip"},
	}

	for _, testCase := range testTable {
		resultStrs, err := Uniq(testCase.input, testCase.opts)
		assert.Equal(t, testCase.expected, resultStrs, fmt.Sprintf("Incorrect result. Expect %v, type of invalid argument: %v, got %v", testCase.expected, testCase.errMsg, resultStrs))
		assert.NotNil(t, err, "Error return value is nil, but expected some error")
	}
}
