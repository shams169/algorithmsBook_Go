package main

import (
	"fmt"
	"math"
)

func selectionSort(A []int) []int {
	B := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		smallest := A[i]
		ind := i
		for j := 0; j < len(A); j++ {
			if A[j] < smallest {
				smallest = A[j]
				ind = j
			}
		}
		B[i] = smallest
		A[ind] = math.MaxInt
	}

	fmt.Println(B)
	return B

}

func selectionSortInplace(A []int) []int {
	for i := 0; i < len(A); i++ {
		sm := A[i]
		ind := i

		for j := i + 1; j < len(A); j++ {
			if A[j] < sm {
				sm = A[j]
				ind = j
			}
		}
		temp := A[i]
		A[i] = A[ind]
		A[ind] = temp
	}
	fmt.Println(A)
	return A
}

func main() {
	A := []int{7, 4, 3, 5, 2, 1, 6}
	//selectionSort(A)
	selectionSortInplace(A)
}
