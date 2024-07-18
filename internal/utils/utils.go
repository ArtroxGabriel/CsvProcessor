package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadCSVFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	aux := []string{""}

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if err != nil || len(line) <= 0 {
			break
		}

		aux = append(
			aux,
			string(line),
		)

	}

	return aux[1:], nil
}

func MapSelectedColumns(header, selectedColumnsString string) (map[int]string, error) {
	hashMap := make(map[int]string)
	selectedColumns := strings.Split(selectedColumnsString, ",")

	for idx := range selectedColumns {
		if !strings.Contains(header, selectedColumns[idx]) {
			return nil, fmt.Errorf("Header '%s' not found in CSV file/string", selectedColumns[idx])
		}
	}
	columns := strings.Split(header, ",")

	if len(selectedColumnsString) == 0 {
		for idx, coluna := range columns {
			hashMap[idx] = coluna
		}
	} else {
		selectedCol := make(map[string]struct{})
		for _, nome := range selectedColumns {
			selectedCol[nome] = struct{}{}
		}
		for idx, coluna := range columns {
			if _, found := selectedCol[coluna]; found {
				hashMap[idx] = coluna
			}
		}
	}
	return hashMap, nil
}
