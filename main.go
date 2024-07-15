package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	usage := "Usage go run . [(a,b)(c,d)] [(e,f)(g,h)]"
	if len(os.Args) != 3 {
		fmt.Println("Error: too many arguments")
		fmt.Println(usage)
		return
	}

	mat1 := os.Args[1]
	mat2 := os.Args[2]

	a, b, c, d := MatrixStr(mat1)
	e, f, g, h := MatrixStr(mat2)
	fmt.Printf("| %3d %3v | *| %3v %3v | \n", a, b, e, f)
	fmt.Printf("| %3v %3v |  | %3v %3v | \n", c, d, g, h)
	i, j, k, l := MultMatrix(a, b, c, d, e, f, g, h)
	fmt.Println("Answer:")
	fmt.Printf("| %5v %5v |\n| %5v %5v | \n", i, j, k, l)
}

// MultMatrix() returns the results of the 2 by 2 matrices multiplication
func MultMatrix(a, b, c, d, e, f, g, h int64) (int64, int64, int64, int64) {
	i := (a * e) + (b * g)
	j := (a * f) + (b * h)
	k := (c * e) + (d * g)
	l := (c * f) + (d * h)
	return i, j, k, l

}

// MatrixStr() takes in a string then returns the matrix values in string form
func MatrixStr(s string) (int64, int64, int64, int64) {
	matSlice := []int64{}
	a := ""

	for _, c := range s {
		if (c >= '0' && c <= '9') || c == '-' {
			a += string(c)
		}
		if !((c >= '0' && c <= '9') || c == '-') && a != "" {
			num, _ := strconv.ParseInt(a, 10, 32)
			matSlice = append(matSlice, num)
			a = ""
		}
	}

	if a != "" {
		num, _ := strconv.ParseInt(a, 10, 32)
		matSlice = append(matSlice, num)
	}

	if len(matSlice) != 4 {
		fmt.Println("Error: innput your matrix using the correct format as shown below:")
		fmt.Println("Usage go run . [(a,b)(c,d)] [(e,f)(g,h)]")
		fmt.Println("Usage go run . replace the letters with numbers")
		os.Exit(0)
	}
	return matSlice[0], matSlice[1], matSlice[2], matSlice[3]
}
