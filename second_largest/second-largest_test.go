package main

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := []int{}
		l := int(math.Pow(2, float64(1+rand.Intn(5))))
		// fmt.Println("length:", l)
		for j := int(0); j < l; j++ {
			n = append(n, int(rand.Intn(100)))
		}
		run(n)
	}
}

func TestSecondLargest(t *testing.T) {

	var testCases = []struct {
		testCase []int
		want     int
	}{
		{
			testCase: []int{1, 2},
			want:     1,
		},
		{
			testCase: []int{2, 1},
			want:     1,
		},
		{
			testCase: []int{59, 9, 77, 76, 0, 19, 22, 5, 42, 38, 37, 44, 1, 23, 54, 37},
			want:     76,
		},
		{
			testCase: []int{4, 2, 3, 1},
			want:     3,
		},
		{
			testCase: []int{4, 1, 5, 3, 6, 2, 8, 7},
			want:     7,
		},
		{
			testCase: []int{2, 1, 4, 3, 5, 7, 8, 6},
			want:     7,
		},
		{
			testCase: []int{14, 19, 1, 30, 16, 88, 68, 87, 71, 93, 38, 64, 83, 91, 45, 12},
			want:     91,
		},
		{
			testCase: []int{664, 407, 700, 165, 56, 620, 981, 146, 137, 326, 766, 511, 957, 143, 216, 239, 15, 963, 912, 246, 553, 48, 762, 142, 275, 364, 258, 535, 204, 183, 123, 376, 374, 332, 597, 761, 849, 85, 951, 523, 579, 530, 678, 62, 27, 802, 257, 889, 470, 832, 852, 537, 114, 97, 442, 636, 389, 992, 375, 364, 654, 4, 364, 979, 213, 758, 977, 381, 783, 722, 154, 615, 456, 127, 179, 851, 410, 26, 2, 502, 400, 919, 750, 36, 833, 109, 408, 200, 466, 555, 557, 473, 497, 940, 755, 815, 612, 735, 315, 956, 147, 505, 80, 113, 321, 233, 928, 39, 833, 289, 244, 661, 640, 226, 32, 899, 549, 384, 816, 296, 435, 162, 190, 408, 133, 899, 2, 699},
			want:     981,
		},
		{
			testCase: []int{49, 52, 92, 41, 96, 16, 80, 68, 33, 47, 1, 23, 11, 4, 89, 56, 17, 93, 82, 67, 51, 15, 97, 91, 17, 8, 38, 38, 41, 96, 94, 8, 26, 22, 13, 16, 17, 11, 51, 25, 5, 55, 73, 6, 45, 85, 81, 30, 76, 69, 77, 66, 87, 96, 95, 69, 35, 68, 11, 5, 57, 57, 92, 67, 83, 19, 95, 72, 18, 7, 9, 93, 78, 68, 18, 18, 60, 61, 41, 97, 17, 38, 50, 79, 44, 98, 27, 63, 34, 16, 51, 40, 58, 0, 22, 59, 63, 94, 71, 45, 70, 34, 90, 85, 49, 64, 79, 13, 62, 52, 73, 55, 62, 79, 93, 16, 71, 54, 14, 3, 13, 15, 50, 14, 16, 15, 53, 35},
			want:     97,
		},
	}

	for _, tt := range testCases {
		currentTest := make([]int, len(tt.testCase))
		copy(currentTest, tt.testCase)
		s, _ := SecondLargestB(currentTest)
		if s != tt.want {
			t.Errorf("SecondLargest(%v) = %d; want %d", tt.testCase, s, tt.want)
		}
	}
}
