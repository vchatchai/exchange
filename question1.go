package main

import "fmt"

func Arrange(repeat int, operations [][]int) []int {

	var ball = []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < repeat; i++ {
		for _, operation := range operations {
			A := operation[0]
			B := operation[1]
			A--
			B--
			ball[A], ball[B] = ball[B], ball[A]
		}
	}

	return ball
}

func main() {

	K := 0
	// /**
	N := 0
	fmt.Println("INPUT")
	fmt.Scanf("%d %d\n", &N, &K)
	var operations = [][]int{}
	for i := 0; i < N; i++ {
		A := 0
		B := 0
		fmt.Scanf("%d %d\n", &A, &B)

		if A > 8 || B > 8 || A < 1 || B < 1 {
			fmt.Println("Please insert value between 1 - 8")
			i--
		} else {
			operations = append(operations, []int{A, B})
		}

	}
	// */
	/*
		K = 1000000000
		operations := [][]int{{1, 3}, {6, 8}, {3, 5}, {2, 6}, {3, 7}, {3, 4}, {4, 7}, {2, 4}, {1, 3}, {2, 7}, {2, 7}, {2, 4}, {6, 7}, {1, 7}, {3, 4}, {1, 6}}
	*/
	result := Arrange(K, operations)
	fmt.Printf("%v", result)
	// result 1, 8, 3, 4, 5, 2, 7, 6

}
