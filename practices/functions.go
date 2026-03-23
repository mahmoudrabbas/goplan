package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}

	fmt.Println(sum(arr...))

	// fmt.Println(arr)

	// arr = transform(&arr, double)
	// fmt.Println(arr)

}

func double(v int) int {
	return v * 2
}

func trible(v int) int {
	return v * 3
}

func transform(arr *[]int, tr func(v int) int) []int {
	for i, v := range *arr {
		(*arr)[i] = tr(v)
	}

	return *arr
}

func sum(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}
