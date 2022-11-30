package main

func Add(number []int) int {

	sum := 0
	// for i := 0; i < 6; i++ {
	// 	sum += number[i]
	// }
	// return sum

	for _, numbers := range number {
		sum += numbers
	}
	return sum
}
