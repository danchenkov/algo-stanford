package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Inversions int = 0

func readLines(path string) ([]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nextInt, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, nextInt)
	}
	return lines, scanner.Err()
}

func main() {
	// original := []int{2, 5, -2, 6, -3, 8, 0, -7, -9, 45, 3, 32, -16, 1, 9, 81, 7, 4}
	// original := []int64{1, 3, 5, 2, 4, 6}
	original, _ := readLines("IntegerArray.txt")

	a := make([]int64, 0, len(original))
	b := make([]int64, 0, len(original))

	// fmt.Printf("MERGE SORT ALGORITHM\n====================\nORIGINAL ARRAY: %v\n", original)
	// Splits
	a = original[:len(original)/2]
	// Split
	b = original[len(original)/2:]
	// fmt.Printf("A: %v\n", a)
	// fmt.Printf("B: %v\n", b)
	// Merge
	MergeSort(a, b)
	// fmt.Printf("SORTED ARRAY:   %v\n", sorted_array)
	fmt.Printf("INVERSIONS: %d\n", Inversions)
}

func MergeSort(a, b []int64) []int64 {
	c := make([]int64, 0, len(a)+len(b))
	// Base case
	if len(a) == 0 {
		copy(c, b)
		// fmt.Printf("%v\n", c)
		return c
	}
	// Base case
	if len(b) == 0 {
		copy(c, a)
		// fmt.Printf("%v\n", c)
		return c
	}
	// Recursion
	if len(a) > 1 {
		a = MergeSort(a[:len(a)/2], a[len(a)/2:])
	}
	if len(b) > 1 {
		b = MergeSort(b[:len(b)/2], b[len(b)/2:])
	}
	// Actual Merge
	for {
		if a[0] < b[0] {
			c = append(c, a[0])
			// fmt.Printf("%v <= %v <= %v %v\n", c, a[0], a, b)
			a = a[1:]
			if len(a) == 0 {
				// fmt.Printf("%v <= ", c)
				c = append(c, b...)
				// fmt.Printf("%v\n%v\n", b, c)
				return c
			}
		} else {
			// fmt.Print(".")
			Inversions += len(a)
			c = append(c, b[0])
			// fmt.Printf("%v <= %v <= %v %v\n", c, b[0], b, a)
			b = b[1:]
			if len(b) == 0 {
				// fmt.Printf("%v <= ", c)
				c = append(c, a...)
				// fmt.Printf("%v\n%v\n", a, c)
				return c
			}
		}
	}
}
