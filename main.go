package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"./uniq"
)

func readStrSlice(reader io.Reader) ([]string, error) {
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(input), "\n"), nil
}

func writeStrSlice(writer io.Writer, strs []string) error {
	_, err := writer.Write([]byte(strings.Join(strs, "\n")))
	return err
}

func parseArgs() (opts uniq.Options, finPath, foutPath string) {
	countPtr := flag.Bool(string(uniq.Count), false, "a bool, prefix lines by the number of occurrences")
	repeatedPtr := flag.Bool(string(uniq.Repeated), false, "a bool, only print duplicate lines, one for each group")
	uniquePtr := flag.Bool(string(uniq.Once), false, "a bool, only print unique lines")
	ignoreCasePtr := flag.Bool(string(uniq.IgnoreCase), false, "a bool, ignore-case, ignore differences in case when comparing")
	skipFieldsPtr := flag.Int(string(uniq.SkipFields), 0, "avoid comparing the first N fields")
	skipCharsPtr := flag.Int(string(uniq.SkipChars), 0, "avoid comparing the first N characters")

	flag.Parse()

	opts.SkipFields = *skipFieldsPtr
	opts.SkipChars = *skipCharsPtr
	opts.IgnoreCase = *ignoreCasePtr

	switch {
	case *countPtr:
		opts.OutputFormat = uniq.Count
	case *repeatedPtr:
		opts.OutputFormat = uniq.Repeated
	case *uniquePtr:
		opts.OutputFormat = uniq.Once
	}

	if len(flag.Args()) >= 1 {
		finPath = flag.Args()[0]
		if len(flag.Args()) >= 2 {
			foutPath = flag.Args()[1]
		}
	}

	return opts, finPath, foutPath
}

func main() {

	opts, finPath, foutPath := parseArgs()

	var reader io.Reader = os.Stdin
	var writer io.Writer = os.Stdout

	if finPath != "" {
		inputFile, err := os.Open(finPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer inputFile.Close()
		reader = inputFile
	}
	if foutPath != "" {
		outputFile, err := os.Create(foutPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer outputFile.Close()
		writer = outputFile

	}

	strsToPocess, err := readStrSlice(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	result_strings, err := uniq.Uniq(strsToPocess, opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writeStrSlice(writer, result_strings)
	if err != nil {
		fmt.Println(err)
	}
}
