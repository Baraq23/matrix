# Matrix Multiplication in Go

This project provides a command-line tool to multiply two 2x2 matrices. The matrices are provided as command-line arguments in a specific format, and the program outputs the result of the multiplication.

## Usage

To run the program, use the following command:

```sh
go run . "[(a,b)(c,d)]" "[(e,f)(g,h)]"

Replace a, b, c, d, e, f, g, h with actual integer values.
Example

sh

go run . "[(1,2)(3,4)]" "[(5,6)(7,8)]"

Output:

makefile

|   1   2 | *|   5   6 | 
|   3   4 |  |   7   8 | 
Answer:
|    19    22 |
|    43    50 | 

Functions
MultMatrix

This function takes eight integers representing two 2x2 matrices and returns four integers representing the result of their multiplication.

go

func MultMatrix(a, b, c, d, e, f, g, h int64) (int64, int64, int64, int64)

MatrixStr

This function takes a string representing a 2x2 matrix in the format "[(a,b)(c,d)]" and returns four integers representing the matrix elements.

go

func MatrixStr(s string) (int64, int64, int64, int64)

Testing

To run the tests for this project, use the following command:

sh

go test

The tests cover both MatrixStr and MultMatrix functions to ensure they work correctly with various input values.
Error Handling

If the input format is incorrect, the program will display an error message and exit. Ensure that the matrices are provided in the correct format as shown in the usage example.
Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue to improve the code or add new features.
License

This project is licensed under the MIT License. See the LICENSE file for details.

go


This `README.md` file provides a comprehensive overview of your project, including usage instructions, function descriptions, testing information, error handling, contributing guidelines, and licensing information.

