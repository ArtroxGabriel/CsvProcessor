#include "libcsv.h"
#include <stdio.h>

int main(void) {
  const char csv[] = "header1,header2,header3\n1,2,3\n4,5,6\n7,8,9";
  processCsv(csv, "header1,header3", "header1>1\nheader3<8");

  printf("\n\n");
  const char csv_file[] = "data.csv";
  processCsvFile(csv_file, "col1,col3,col4,col7", "col1>l1c1\ncol3>l1c3");

  return 0;
}
