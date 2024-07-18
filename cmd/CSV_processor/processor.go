package csvprocessor

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/ArtroxGabriel/CSVProcessor/internal/filter"
	"github.com/ArtroxGabriel/CSVProcessor/internal/utils"
)

var logStderr = log.New(os.Stderr, "", 0)

type processor struct {
	filters         map[string]filter.Filter
	selectedColumns map[int]string
	csvData         []string
}

func NewProcessor() CSVProcessor {
	return &processor{}
}

func (proc *processor) ProcessCSV(csvSlice []string, selectedColumns, rowFilterDefinitions string) {
	proc.csvData = csvSlice[1:]

	// get selectedColumns
	selectedColsMap, err := utils.MapSelectedColumns(csvSlice[0], selectedColumns)
	if err != nil {
		logStderr.Println(err)
		return
	}
	proc.selectedColumns = selectedColsMap

	// get filters
	filtersHash, err := filter.GetFilters(csvSlice[0], rowFilterDefinitions)
	if err != nil {
		logStderr.Println(err)
		return
	}
	proc.filters = filtersHash

	proc.filterData()

	proc.stdout(selectedColumns)
}

func (proc *processor) filterData() {
	filteredData := make([][]string, 0, len(proc.csvData))
	indexOrdered := make([]int, 0, len(proc.selectedColumns))
	for idx := range proc.selectedColumns {
		indexOrdered = append(indexOrdered, idx)
	}
	sort.Ints(indexOrdered)

	for _, row := range proc.csvData {
		values := strings.Split(row, ",")
		filteredRow := make([]string, 0, len(proc.selectedColumns))

		for _, idx := range indexOrdered {
			colName := proc.selectedColumns[idx]

			if filter, exists := proc.filters[colName]; !exists || filter.Filtrar(values[idx]) {
				filteredRow = append(filteredRow, values[idx])
			} else {
				break
			}
		}

		if len(filteredRow) == len(proc.selectedColumns) {
			filteredData = append(filteredData, filteredRow)
		}
	}

	proc.csvData = filteredDataToStrings(filteredData)
}

func filteredDataToStrings(filteredData [][]string) []string {
	result := make([]string, len(filteredData))
	for i, row := range filteredData {
		result[i] = strings.Join(row, ",")
	}
	return result
}

func (proc *processor) stdout(selectedColumns string) {
	if selectedColumns == "" {
		stringOut := make([]string, 0)
		for idx := 0; idx < len(proc.selectedColumns); idx++ {
			stringOut = append(stringOut, proc.selectedColumns[idx])
		}
		fmt.Println(strings.Join(stringOut, ","))
	} else {
		fmt.Println(selectedColumns)
	}

	std := strings.Join(proc.csvData, "\n")
	fmt.Println(std)
}
