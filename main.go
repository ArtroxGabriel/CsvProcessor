package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unsafe"

	csvprocessor "github.com/ArtroxGabriel/CSVProcessor/cmd/CSV_processor"
	"github.com/ArtroxGabriel/CSVProcessor/internal/utils"
)

var logStderr = log.New(os.Stderr, "", 0)

//export processCsvGo
func processCsvGo(csv, selectedColumns, rowFilterDefinitions *C.char) {
	processor := csvprocessor.NewProcessor()

	goCsv := C.GoString(csv)
	goSelectedColumns := strings.TrimSpace(C.GoString(selectedColumns))
	goRowFilterDefinitions := strings.TrimSpace(C.GoString(rowFilterDefinitions))

	csvSlice := strings.Split(goCsv, "\n")
	processor.ProcessCSV(
		csvSlice,
		goSelectedColumns,
		goRowFilterDefinitions,
	)
}

//export processCsvFileGo
func processCsvFileGo(csvFilePath, selectedColumns, rowFilterDefinitions *C.char) {
	processor := csvprocessor.NewProcessor()

	goCsvFilePath := C.GoString(csvFilePath)
	goSelectedColumns := strings.TrimSpace(C.GoString(selectedColumns))
	goRowFilterDefinitions := strings.TrimSpace(C.GoString(rowFilterDefinitions))

	csvSlices, err := utils.ReadCSVFile(goCsvFilePath)
	if err != nil {
		logStderr.Print(err)
		return
	}

	processor.ProcessCSV(
		csvSlices,
		goSelectedColumns,
		goRowFilterDefinitions,
	)
}

func main() {
	csvData := "header1,header2,header3\n1,2,3\n4,5,6\n7,8,9"
	selectedColumns := "header1,header3"
	rowFilterDefinitions := "header1>1\nheader3<8"

	// Test processCsvGo (inline CSV data)
	fmt.Println("Testing processCsvGo:")
	cCsvData := C.CString(csvData)
	cSelectedColumns := C.CString(selectedColumns)
	cRowFilterDefinitions := C.CString(rowFilterDefinitions)

	processCsvGo(cCsvData, cSelectedColumns, cRowFilterDefinitions)

	C.free(unsafe.Pointer(cCsvData))

	fmt.Println("\nTesting processCsvFileGo (with data.csv):")
	cCsvFilePath := C.CString("data.csv")
	cSelectedColumns = C.CString("col1,col3,col4,col7")
	cRowFilterDefinitions = C.CString("col1>l1c1\ncol3>l1c3")

	processCsvFileGo(cCsvFilePath, cSelectedColumns, cRowFilterDefinitions)

	C.free(unsafe.Pointer(cCsvFilePath))
	C.free(unsafe.Pointer(cSelectedColumns))
	C.free(unsafe.Pointer(cRowFilterDefinitions))
}
