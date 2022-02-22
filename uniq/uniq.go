package uniq

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	UniqDefault byte = 0
	Count       byte = 'c'
	Repeated    byte = 'd'
	Once        byte = 'u'
	IgnoreCase  byte = 'i'
	SkipFields  byte = 'f'
	SkipChars   byte = 's'
)

type Options struct {
	OutputFormat byte
	SkipFields   int
	SkipChars    int
	IgnoreCase   bool
}

func compareStrings(l, r string) bool {
	return l == r
}

func skipFieldsChars(str string, numSkipFields int, numSkipChars int) string {
	if numSkipFields == 0 {
		if numSkipChars > len(str) {
			return ""
		}
		return string(str[numSkipChars:])
	}

	beginIndex := 0
	inField := false
	for index, currRune := range " " + str {
		if numSkipFields < 0 {
			break
		}
		if !unicode.IsSpace(currRune) && !inField {
			inField = true
			beginIndex = index
			numSkipFields--
		}
		if unicode.IsSpace(currRune) {
			inField = false
		}
	}
	if numSkipFields >= 0 || beginIndex+numSkipChars-1 > len(str) {
		return ""
	}
	return str[beginIndex-1+numSkipChars:]
}

func copySkipingFieldsCharsInSlice(strs []string, skipFields int, skipChars int) (resultStrs []string, err error) {
	if strs == nil {
		return nil, fmt.Errorf("nil slice passed")
	}
	if skipFields < 0 || skipChars < 0 {
		return nil, fmt.Errorf("number of fields/characters to skip cannot be negative")
	}
	resultStrs = make([]string, len(strs))
	if skipFields != 0 || skipChars != 0 {
		for index, str := range strs {
			resultStrs[index] = skipFieldsChars(str, skipFields, skipChars)
		}
	} else {
		copy(resultStrs, strs)
	}
	return resultStrs, nil
}

func uniq(strs []string, compare func(string, string) bool, skipFields int, skipChars int) (resultStrs []string, err error) {
	if strs == nil || compare == nil {
		return nil, fmt.Errorf("nil argument passed")
	}

	strsSkipped, err := copySkipingFieldsCharsInSlice(strs, skipFields, skipChars)
	if err != nil {
		return nil, err
	}

	prevIndex := 0
	var index int
	var currStr string
	for index, currStr = range strsSkipped {
		if compare(currStr, strsSkipped[prevIndex]) {
			continue
		}

		resultStrs = append(resultStrs, strs[prevIndex])
		prevIndex = index
	}
	resultStrs = append(resultStrs, strs[index])
	return resultStrs, nil
}

func uniqCountConditional(strs []string, compare func(string, string) bool, countConditional func(int) bool, skipFields int, skipChars int) (repeatedStrs []string, err error) {
	if strs == nil || compare == nil {
		return nil, fmt.Errorf("nil argument passed")
	}

	strsSkipped, err := copySkipingFieldsCharsInSlice(strs, skipFields, skipChars)
	if err != nil {
		return nil, err
	}

	count := 0
	prevIndex := 0
	var currStr string
	var index int
	for index, currStr = range strsSkipped {
		if compare(currStr, strsSkipped[prevIndex]) {
			count++
			continue
		}
		if countConditional(count) {
			repeatedStrs = append(repeatedStrs, strs[prevIndex])
		}
		prevIndex = index
		count = 1
	}
	if countConditional(count) {
		repeatedStrs = append(repeatedStrs, strs[index])
	}
	return repeatedStrs, nil
}

func uniqCount(strs []string, compare func(string, string) bool, skipFields int, skipChars int) (uniqueStrs []string, err error) {
	if strs == nil || compare == nil {
		return nil, fmt.Errorf("nil argument passed")
	}

	strsSkipped, err := copySkipingFieldsCharsInSlice(strs, skipFields, skipChars)
	if err != nil {
		return nil, err
	}

	count := 0
	prevIndex := 0
	var currStr string
	var index int

	for index, currStr = range strsSkipped {
		if compare(currStr, strsSkipped[prevIndex]) {
			count++
			continue
		}
		uniqueStrs = append(uniqueStrs, strconv.Itoa(count)+" "+strs[prevIndex])
		prevIndex = index
		count = 1
	}
	uniqueStrs = append(uniqueStrs, strconv.Itoa(count)+" "+strs[index])
	return uniqueStrs, nil
}

func Uniq(strs []string, opts Options) ([]string, error) {
	if strs == nil {
		return nil, fmt.Errorf("nil slice passed")
	}

	compareFunc := compareStrings
	if opts.IgnoreCase {
		compareFunc = strings.EqualFold
	}

	switch opts.OutputFormat {
	case UniqDefault:
		return uniq(strs, compareFunc, opts.SkipFields, opts.SkipChars)
	case Once:
		return uniqCountConditional(strs, compareFunc, func(count int) bool { return count == 1 }, opts.SkipFields, opts.SkipChars)
	case Count:
		return uniqCount(strs, compareFunc, opts.SkipFields, opts.SkipChars)
	case Repeated:
		return uniqCountConditional(strs, compareFunc, func(count int) bool { return count > 1 }, opts.SkipFields, opts.SkipChars)
	default:
		return nil, fmt.Errorf("unknown option %c", opts.OutputFormat)
	}
}
