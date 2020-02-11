package main

import (
	"errors"
	"log"
	"math/bits"
)

// You are given as input an unsorted array of n distinct numbers, where n is a power of 2.
// Give an algorithm that identifies the second-largest number in the array, and that uses at most `n+log(n)-2` comparisons.
func main() {
	// run(n)
	// n := []int{2, 1}
	// n := []int{2, 4, 3, 1}
	// n := []int{2, 4, 6, 3, 8, 1, 5, 7}
	n := []int{664, 407, 700, 165, 56, 620, 981, 146, 137, 326, 766, 511, 957, 143, 216, 239, 15, 963, 912, 246, 553, 48, 762, 142, 275, 364, 258, 535, 204, 183, 123, 376, 374, 332, 597, 761, 849, 85, 951, 523, 579, 530, 678, 62, 27, 802, 257, 889, 470, 832, 852, 537, 114, 97, 442, 636, 389, 992, 375, 364, 654, 4, 364, 979, 213, 758, 977, 381, 783, 722, 154, 615, 456, 127, 179, 851, 410, 26, 2, 502, 400, 919, 750, 36, 833, 109, 408, 200, 466, 555, 557, 473, 497, 940, 755, 815, 612, 735, 315, 956, 147, 505, 80, 113, 321, 233, 928, 39, 833, 289, 244, 661, 640, 226, 32, 899, 549, 384, 816, 296, 435, 162, 190, 408, 133, 899, 2, 699}
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
	SecondLargestB(n)
}

// pair implementation with running time O(2*n)
func SecondLargestA(n []int) (int, int) {
	var largest, second int
	var counter int
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
			counter++
		}
		counter++
	}
	return second, counter
}

// pair implementation with running time O(n + log(n) - 2)
func SecondLargestB(n []int) (int, int) {
	var counter int
	losses := map[int][]int{}
	for len(n) > 1 {
		l := len(n) / 2
		for i := 0; i < l; i++ {
			// counter++
			if n[i] > n[l+i] {
				losses[n[i]] = append(losses[n[i]], n[l+i])
			} else {
				losses[n[l+i]] = append(losses[n[l+i]], n[i])
				n[i] = n[l+i]
			}
			counter++
		}
		n = n[0:l]
		// counter++
	}
	f := n[0]
	s, loss := losses[f][0], losses[f][1:]
	for _, c := range loss {
		if c > s {
			counter++
			s = c
		}
	}
	return s, counter
}
