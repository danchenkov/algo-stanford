package main

import (
	"errors"
	"fmt"
	"log"
	"math/bits"
)

// You are given as input an unsorted array of n distinct numbers, where n is a power of 2.
// Give an algorithm that identifies the second-largest number in the array, and that uses at most `n+log(n)-2` comparisons.
func main() {
	// n := []int{2, 1}
	// run(n)
	n := []int{2, 4, 3, 1}
	run(n)
	// rand.Seed(time.Now().UTC().UnixNano())
	// rand.Seed(1580956330)
	// for i := 0; i < 5; i++ {
	// 	n := []int{}
	// 	l := int(math.Pow(2, float64(1+rand.Intn(7))))
	// 	// fmt.Println("length:", l)
	// 	for j := 0; j < l; j++ {
	// 		n = append(n, rand.Intn(100))
	// 	}
	// 	run(n)
	// }
}

func run(n []int) {
	original_n := make([]int, len(n))
	copy(original_n, n)
	if len(n) < 2 {
		log.Fatal(errors.New("Insufficient length of array"))
	}
	if bits.OnesCount(uint(len(n))) != 1 {
		log.Fatal(errors.New("Array length is not of power of two"))
	}
	s := SecondLargestB(n)
	fmt.Println(s)
}

func log2_32(value uint32) uint32 {
	tab32 := [32]uint32{0, 9, 1, 10, 13, 21, 2, 29, 11, 14, 16, 18, 22, 25, 3, 30, 8, 12, 20, 28, 15, 17, 24, 7, 19, 27, 23, 6, 26, 5, 4, 31}
	value |= value >> 1
	value |= value >> 2
	value |= value >> 4
	value |= value >> 8
	value |= value >> 16
	return tab32[uint32(value*0x07C4ACDD)>>27]
}

// pair implementation with running time O(2*n)
func SecondLargestA(n []int) int {
	var largest, second int
	if n[0] > n[1] {
		largest = n[0]
		second = n[1]
	} else {
		largest = n[1]
		second = n[0]
	}
	for i := 2; i < len(n); i++ {
		if n[i] > largest {
			second = largest
			largest = n[i]
		} else {
			if n[i] >= second {
				second = n[i]
			}
		}
	}
	return second
}

func iPow(a int, b int) int {
	var result int = 1
	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}
	return result
}

// pair implementation with running time O(n + log(n) - 2)
func SecondLargestB(n []int) int {
	losses := map[int][]int{}

	// fmt.Printf("N: %v\n", n)
	// fmt.Printf("len(n): %d\n", len(n))
	// for level := 0; (2 >> level) < len(n); level++ {
	// 	fmt.Printf("Level: %d\n", level)
	// 	for i := 0; i+2<<level < len(n); i += iPow(2, level) {
	// 		fmt.Printf("i: %d, iterator: %d\n", i, 2>>level)
	// 		fmt.Printf("n[%d]=%d vs n[%d]=%d\n", i, n[i], i+2<<level, n[i+2<<level])
	// 		if n[i] > n[i+2<<level] {
	// 			losses[i] = append(losses[i], n[i+2<<level])
	// 		} else {
	// 			n[i] = n[i+2<<level]
	// 		}
	// 	}
	// }
	// fmt.Printf("N: %v\n", n)
	for len(n) > 1 {
		l := len(n) / 2
		for i := 0; i < l; i++ {
			if n[i] > n[l+i] {
				losses[n[i]] = append(losses[n[i]], n[l+i])
			} else {
				losses[n[l+i]] = append(losses[n[l+i]], n[i])
				n[i] = n[l+i]
			}
			// fmt.Println("Losers: ", losses)
		}
		n = n[0:l]
		// fmt.Printf("N: %v\n", n)
	}
	f := n[0]
	// fmt.Println("Winner: ", f)
	// fmt.Println("Losers: ", losses[f])
	s, loss := losses[f][0], losses[f][1:]
	// fmt.Println("Candidate: ", s)
	// fmt.Println(" or one of ", loss)
	// fmt.Printf("len(loss): %d", len(loss))
	for _, c := range loss {
		if c > s {
			s = c
		}
	}
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	return s
}

// func SecondLargestB(n []int) (s int) {
// 	var half int
// 	var level int
// 	var i int
// 	// maxlevel := int(math.Log2(float64(len(n))))
// 	maxlevel := log2_32(int(len(n)))
// 	losses := make([]int, 0, maxlevel)
// 	// fmt.Printf("   I: %v\n", n)
// 	for level = 0; level < maxlevel-1; level++ {
// 		half = int(len(n) / (2 << level))
// 		for i = 0; i < half; i++ {
// 			if n[i+half] > n[i] {
// 				losses = append(losses, n[i])
// 				n[i] = n[i+half]
// 			} else {
// 				losses = append(losses, n[i+half])
// 			}
// 		}
// 		// fmt.Printf("%4d: %v\n", level, n[0:len(n)/(2<<level)])
// 	}
// 	// fmt.Printf("LOSSES: %v\n", losses)
// 	if n[1] > n[0] {
// 		s = n[0]
// 	} else {
// 		s = n[1]
// 	}
// 	for i = 0; i < int(len(losses)); i++ {
// 		if losses[i] > s {
// 			s = losses[i]
// 		}
// 	}
// 	return s
// }
