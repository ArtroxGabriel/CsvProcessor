# Processador de CSV

## Objective

Implement a library that processes CSV files, applying filters and selecting columns as specified. The solution must be able to integrate with the C interface defined below.

```c
/**
 * Process the CSV data by applying filters and selecting columns.
 *
 * @param csv The CSV data to be processed.
 * @param selectedColumns The columns to be selected from the CSV data.
 * @param rowFilterDefinitions The filters to be applied to the CSV data.
 *
 * @return void
 */
void processCsv(const char[], const char[], const char[]);

/**
 * Process the CSV file by applying filters and selecting columns.
 *
 * @param csvFilePath The path to the CSV file to be processed.
 * @param selectedColumns The columns to be selected from the CSV data.
 * @param rowFilterDefinitions The filters to be applied to the CSV data.
 *
 * @return void
 */
void processCsvFile(const char[], const char[], const char[]);
```

processCsv

- csv: A string with the CSV data, where each line represents a record, and the columns are separated by commas.
  - Example: `"header1,header2,header3\n1,2,3\n4,5,6"`
- selectedColumns: A string where the names of the columns to be selected are separated
  - Example: `"header1,header3"`
- rowFilterDefinitions:  A string where each filter is defined on a new line, in the format `header(comparator)value`.
  - Example: `"header1>1\nheader2=2\nheader3<6"`

processCsvFile

- csvFilePath: A string with the path to the CSV file.
  - Example: `"path/to/csv_file.csv"`
- selectedColumns: A string where the names of the columns to be selected are separated
  - Example: `"header1,header3"`
- rowFilterDefinitions:  A string where each filter is defined on a new line, in the format `header(comparator)value`.
  - Example: `"header1>1\nheader2=2\nheader3<6"`

Example:

```c
const char csv[] = "header1,header2,header3,header4\n1,2,3,4\n5,6,7,8\n9,10,11,12";
processCsv(csv, "header1,header3,header4", "header1>1\nheader3<10");

// output:
// header1,header3,header4
// 5,7,8


const char csv_file[] = "path/to/csv_file.csv";
processCsvFile(csv_file, "header1,header3,header4", "header1>1\nheader3<10");

// output:
// header1,header3,header4
// 5,7,8
```

## Mandatory Features and Requirements

For the examples below, always consider the following CSV:
```csv
header1,header2,header3
1,2,3
4,5,6
7,8,9
```

- **The processed CSV must be written to stdout**
- **The first line of the CSV will always be a header**
- **The processed CSV should include the header, considering the column selection**
- **Your implementation must handle CSVs with arbitrary amounts of characters**
- **Your implementation must handle CSVs where columns have arbitrary amounts of characters**
- **An empty column selection string is equivalent to selecting all columns**
- **At a minimum, the candidate must implement filters for greater than (>), less than (<), and equal to (=)**
- **The column selection and filter strings will always be in the same order as they appear in the CSV**


  Example:

  - `"header1,header3"` ou `"header1=4\nheader3>3"` &rarr; OK
  - `"header3,header1"` ou `"header3>3\nheader1=4"` &rarr; NÃO OK

- **Commas always delimit a column; quotes have no special interpretation**

  Example: In the CSV below, the name of the first column is  `'hea"der1'`

  ```csv
  hea"der1,header2,header3
  1,2,3
  ```

- **Only rows that match all filters should be selected**

  Example: Applying the filters `"header1=4\nheader2>3"` and selecting the columns header1 and header3. Only the
  row 4,5,6 `(header1 = 4 AND header2 > 3)` should be selected, as all filter conditions must be met. Filter output below:



  ```csv
  header1,header3
  4,6
  ```

- **Invalid filters or nonexistent columns will never be provided**

  Example:

  - Nonexistent column: `"header4"`
  - Invalid filter: `"header1#2"`

- **No more than one filter per column will be provided**

  Example: If the filter is `"header1=2"`, there will be no other filter for `"header1"` in the same operation.

- **The input CSV will have a maximum of 256 columns**
- **The name of each column is unique**
- **Comparators in filters must follow lexicographical order according to the [strcmp](https://www.man7.org/linux/man-pages/man3/strcmp.3.html) implementation of the stdlibc**
- **The target architecture is x86_64**
- **Executing external processes is not allowed. Your code should not use system calls to execute other programs**
- **Using external libraries to convert the CSV into intermediate data structures is not allowed (libs for lexers and tokenizers can be used)**

## Bonus Features

For all the examples below, always consider the following CSV:

```csv
header1,header2,header3
1,2,3
4,5,6
7,8,9
```

1. **Develop unit tests**
2. **Columns in the selected columns string can be in arbitrary order**

   Example: If the selected columns string is `"header3,header1"`, your implementation should select the columns in this order.

   ```csv
   header1,header2,header3
   1,2,3
   4,5,6
   7,8,9
   ```

  And the selected columns string: `"header3,header1"`, the result should be:

   ```csv
   header1, header3
   1,3
   4,6
   7,9
   ```

3. **Filters can appear in arbitrary order in the filter string**

   If the filters are provided as `"header2>3\nheader1=4"`, your implementation should apply these filters correctly,
   regardless of the order, as shown in the output below:

   ```csv
   header1,header2,header3
   4,5,6
   ```
4. **Nonexistent columns can appear in the selected columns and filters**

If the selected columns string includes "header4" and the CSV does not have a header header4, or if the filters  
  include `"header5=10"` and the CSV does not have a header header5, your implementation should handle these cases
  appropriately by writing to stderr with the message `"Header 'header4'` not found in CSV file/string" or
  `"Header 'header5' not found in CSV file/string"`, respectively, and terminating the execution.

5. **Error handling for invalid filters**

  If an invalid or nonexistent filter is provided, such as `header1#2`, your implementation should handle these 
  cases by writing to stderr with the message `"Invalid filter: 'header1#2'"`.
  
6. **Accept more than one filter per header**

  Allow multiple filters to be applied to the same header, such as `"header1=1\nheader1=4\nheader2>3\nheader3>4"`, 
  and handle these filters appropriately. For this implementation, the candidate should implement the OR logic for 
  filters on the same header. 
  The example filter will be considered as: `(header1=1 OR header1=4) AND header2>3 AND header3>4`, resulting in the output below:

   ```csv
   header1,header2,header3
   4,5,6
   ```

7. **Implement the not equal (!=), greater than or equal to (>=), and less than or equal to (<=) operators**

  Allow filters to use not equal, greater than or equal, and less than or equal operators, such as `"header1!=2\nheader2>=5\nheader3<=6"`, 
  and handle these filters appropriately, resulting in the output below:

   ```csv
   header1,header2,header3
   7,8,9
   ```

---
