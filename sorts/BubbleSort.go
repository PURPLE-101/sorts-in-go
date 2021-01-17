package main

import "fmt"

func bubbleSort(numbers[]int) []int{

	arrayLength := len(numbers)

	for i := 0; i < arrayLength; i++ {
		for j := 1; j < arrayLength-1; j++ {
			if(numbers[j-1] > numbers[j]) {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}

	return numbers
}

func main(){
	var numbers = []int {1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	
	fmt.Println(bubbleSort(numbers))	
}