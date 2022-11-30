package main

func Add(a [5]int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += a[i]
	// }
	// return sum

	for _, num := range a {
		sum += num
	}
	return sum
}
