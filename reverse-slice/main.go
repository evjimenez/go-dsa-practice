package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	rotationSlice(nums, 2)
}

func rotationSlice(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n // Normalizar k

	reverse(nums, 0, n-1) //invertir todo
	reverse(nums, 0, k-1) //invertir primeros k
	reverse(nums, k, n-1) // invertir el resto

	fmt.Println(nums)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
