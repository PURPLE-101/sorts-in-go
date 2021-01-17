package main

import "fmt"

func bubbleSort(numbers[10]int) [10]int{

	arrayLength := len(numbers)

	for i := 0; i < arrayLength; i++ {
		for j := 1; j < arrayLength-1; j++ {
			if(numbers[j-1] > numbers[j]) {
				temp := numbers[j-1]
				numbers[j-1] = numbers[j]
				numbers[j] = temp
			}
		}
	}

	return numbers
}

func main(){
	var numbers = [10]int {1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	
	fmt.Println(bubbleSort(numbers))	
}