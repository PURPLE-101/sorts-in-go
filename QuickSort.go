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

			swaptemp := numbers[i]
			numbers[i] = numbers[j]
            numbers[j] = swaptemp
		}
	}

	swapTemp := numbers[i+1]
	numbers[i+1] = numbers[end]
	numbers[end] = swapTemp

	return i+1
}

func main(){
	var numbers = []int {1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	arrayLength := len(numbers)
	fmt.Println(quickSort(numbers,0, arrayLength-1))
	
	fmt.Println("----------------------")

}