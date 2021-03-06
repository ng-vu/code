package main

import (
	"fmt"
	"testing"
)

func TestSumGCD(t *testing.T) {
	var tests = [][2][]int{
		{
			{113, 31, 23, 47, 87, 17, 53, 41, 67, 61, 13, 7, 83, 37, 101},
			{113, 31, 23, 47, 87, 17, 53, 41, 67, 61, 13, 7, 83, 37, 101},
		},
		{
			{113, 31, 23, 47, 87, 17, 53, 41, 67, 61, 13, 7, 83, 37, 62},
			{113, 62, 31, 23, 47, 87, 17, 53, 41, 67, 61, 13, 7, 83, 37},
		},
		{
			{113, 31, 23, 48, 87, 17, 53, 41, 67, 61, 13, 7, 83, 37, 62},
			{113, 62, 31, 23, 48, 87, 17, 53, 41, 67, 61, 13, 7, 83, 37},
		},
		{
			{114, 31, 23, 48, 87, 18, 53, 41, 67, 61, 13, 7, 83, 37, 62},
			{87, 48, 114, 18, 62, 31, 23, 53, 41, 67, 61, 13, 7, 83, 37},
		},
		{
			{1, 2, 4, 8, 16, 32, 64, 128, 256, 511},
			{511, 256, 128, 64, 32, 16, 8, 4, 2, 1},
		},
		{
			{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			{2, 14, 7, 5, 15, 20, 10, 11, 13, 19, 3, 9, 18, 6, 12, 8, 16, 4, 1, 17},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			A = tt[0]
			main()
			expect(t, tt[1])
		})
	}
}

func expect(t *testing.T, expected []int) {
	expectedMax := sumGCD(expected)
	if max != expectedMax {
		t.Errorf("Expected: %v %v", expectedMax, expected)
	}
}

func sumGCD(list []int) int {
	s := 0
	for i := 1; i < len(list); i++ {
		s += gcd(list[i-1], list[i])
	}
	return s
}
