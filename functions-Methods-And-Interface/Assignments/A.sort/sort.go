package main

import "fmt"

func Swap(slice *[]int, index int) {
	(*slice)[index], (*slice)[index+1] = (*slice)[index+1], (*slice)[index]
}

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(&slice, j)
			}
		}
	}
	return slice
}

func TakeInput(arr []int, err error) []int {
	if err != nil {
		return arr
	}

	if len(arr) == 10 {
		return arr
	}

	var inputInt int
	fmt.Print("Enter number input or word to exit : ")
	n, err := fmt.Scan(&inputInt)
	if n == 1 {
		arr = append(arr, inputInt)
	}
	return TakeInput(arr, err)
}

func main() {
	slice := []int{}
	slice = TakeInput(slice, nil)
	fmt.Println("Input : ", slice)
	slice = BubbleSort(slice)
	fmt.Println("Sorted : ", slice)
}
