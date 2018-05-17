package main

import "fmt"

func nextPermutation(a []int) {
	if len(a) < 2 {
		return
	}
	i, j := len(a)-1, len(a)-2
	var p, q int
	for j >= 0 && a[i] < a[j] {
		i--
		j--
	}

	if j < 0 {
		j = 0
		i = 1
	}

	p = j

	for i < len(a) && a[i] > a[p] {
		i++
	}

	q = i - 1

	//	fmt.Println(p, q)
	tp := a[p]
	a[p] = a[q]

	sl := len(a) - 1 - q

	//	fmt.Println("hehre", sl)
	for i = 0; i < sl+1; i++ {
		//		fmt.Println(i)
		t := a[len(a)-1]
		j = len(a) - 1
		for j > p+i+1 {
			a[j] = a[j-1]
			j--
		}
		a[j] = t
	}

	//	fmt.Println("ooo")
	a[p+sl+1] = tp
	i = p + sl + 2
	j = len(a) - 1
	//	fmt.Println(i, j)
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	fmt.Println(a)
}

func main() {
	nums := []int{1, 1}
	//	fmt.Println(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	nextPermutation(nums)
	//	fmt.Println(nums)
}
