package main

import (
	"fmt"
)

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

func ArrangeOpmized(repeat int, operations [][]int) []int {
	var balls = []int{1, 2, 3, 4, 5, 6, 7, 8}
	indexes := indexSet(operations)
	result := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for c := 0; c < repeat; c++ {
		for i, ball := range balls {
			next := indexes[i]
			result[next] = ball
		}
		copy(balls, result)
	}
	return balls
}

func ArrangeGORoutine(repeat int, operations [][]int) []int {
	var balls = []int{1, 2, 3, 4, 5, 6, 7, 8}
	indexes := indexSet(operations)
	result := []int{1, 2, 3, 4, 5, 6, 7, 8}

	input := make(chan []int, 8)
	output := make(chan []int, 8)

	for c := 0; c < 8; c++ {
		go routine(repeat, indexes, input, output)
	}

	for i, ball := range balls {
		values := []int{i, ball}
		input <- values
	}
	count := 0
	for values := range output {
		result[values[0]] = values[1]
		count++
		if count == 8 {
			close(input)
			close(output)
		}
	}

	return result
}

func routine(repeat int, indexes []int, input <-chan []int, output chan<- []int) {
	for values := range input {
		i := values[0]
		ball := values[1]
		for c := 0; c < repeat; c++ {
			i = indexes[i]
		}
		output <- []int{i, ball}
	}

}

func indexSet(operations [][]int) []int {
	var balls = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var indexes = []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, operation := range operations {
		A := operation[0]
		B := operation[1]
		A--
		B--
		balls[A], balls[B] = balls[B], balls[A]
	}
	for i, ball := range balls {
		indexes[ball-1] = i
	}

	return indexes
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
	result := ArrangeGORoutine(K, operations)
	fmt.Printf("%v", result)
	// result 1, 8, 3, 4, 5, 2, 7, 6

}
