package main

import "fmt"

func quickSort(numbers[]int, begin, end int) []int{
	if(begin < end) {
		var partitionIndex = partition(numbers, begin, end);

        numbers = quickSort(numbers, begin, partitionIndex-1);
        numbers = quickSort(numbers, partitionIndex+1, end);
	}

	return numbers
}

func partition(numbers[]int, begin, end int) int{
	var pivot = numbers[end]
	var i = (begin -1)

	for j := begin; j < end; j++ {
		if(numbers[j] <= pivot) {
			i++
			numbers = swapNumbers(numbers, i, j)
		}
	}

	numbers = swapNumbers(numbers, i+1, end)

	return i+1
}

func swapNumbers(numbers[]int, a, b int) []int{
	swapTemp := numbers[a]
	numbers[a] = numbers[b]
	numbers[b] = swapTemp

	return numbers
}

func main(){
	var numbers = []int {1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	arrayLength := len(numbers)
	fmt.Println(quickSort(numbers,0, arrayLength-1))
	
	fmt.Println("----------------------")
}