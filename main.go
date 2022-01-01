package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	res, err := run()
	fmt.Println(res)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

func run() (string, error) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), int(1e11))
	if !sc.Scan() {
		return "", fmt.Errorf("failed to scan header")
	}
	headers := strings.Split(sc.Text(), ",")

	df := make([][]string, 0, 100)
	for sc.Scan() {
		row := strings.Split(sc.Text(), ",")
		if len(row) != len(headers) {
			return toString(headers, df), fmt.Errorf("number of header columns and body columns is not equal")
		}
		df = append(df, row)
	}
	return toString(headers, df), nil
}

func toString(header []string, df [][]string) string {
	builder := &strings.Builder{}
    for colNum, col := range header {
		builder.WriteString(col)
		if colNum != len(header)-1 {
			builder.WriteString("\t")
		}
    }
	builder.WriteString("\n")
	for rowNum, row := range df {
		for colNum, col := range row {
			builder.WriteString(col)
			if colNum != len(row)-1 {
				builder.WriteString("\t")
			}
		}
		if rowNum != len(df)-1 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
