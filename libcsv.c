#include "libcsv.h"
#include "libcsvprocessor.h"

void processCsv(const char csv[], const char selectedColumns[],
                const char rowFilterDefinitions[]) {
  processCsvGo((char *)csv, (char *)selectedColumns,
               (char *)rowFilterDefinitions);
}

void processCsvFile(const char csvFilePath[], const char selectedColumns[],
                    const char rowFilterDefinitions[]) {
  processCsvFileGo((char *)csvFilePath, (char *)selectedColumns,
                   (char *)rowFilterDefinitions);
}
