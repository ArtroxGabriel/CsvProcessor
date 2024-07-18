package csvprocessor

type CSVProcessor interface {
	// Process the CSV data by applying filters and selecting columns.
	//
	// Parameters:
	//
	//	csvSlices: The CSV data to be processed.
	//	selectedColumns: A comma-separated list of column names to be selected.
	//	rowFilterDefinitions: A JSON string defining the filters to be applied to the CSV rows.
	ProcessCSV(csvSlice []string, selectedColumns, rowFilterDefinitions string)
}
