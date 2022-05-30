package main

import (
	"fmt"
	"time"
)

const NUM_THREADS = 5
const ARRAY_SIZE = 1000000000

func main() {
	var arr []uint64
	var start, end time.Time
	var sum uint64 = 0
	var i uint64 = 0
	arr = make([]uint64, ARRAY_SIZE)

	// Initialize the array and measure the time
	start = time.Now()
	for i = 0; i < ARRAY_SIZE; i++ {
		arr[i] = i + 1
	}
	end = time.Now()
	fmt.Printf("Initialization Took: %d ms\n\n", end.Sub(start)/1000000)

	// Sum the array on one thread and measure the time
	start = time.Now()
	// for i = 0; i < ARRAY_SIZE; i++ {
	// 	sum += arr[i]
	// }
	sum = ArraySum(&arr)
	end = time.Now()
	fmt.Printf("Sum of the Array: %d \n", sum)
	fmt.Printf("Elapsed Time: %d ms\n", end.Sub(start)/1000000)
}

func ArraySum(arr *[]uint64) (sum uint64) {
	// sum = 0
	// arrP, ok := arr.(*[]uint64)
	// if !ok {
	// 	fmt.Printf("failed")
	// 	return
	// }
	
	for i := uint64(0); i < ARRAY_SIZE; i++ {
		sum += (*arr)[i]
	}
	return
}
