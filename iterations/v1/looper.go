package main

import (
	"fmt"
)

func Repeater(num []int) int {

	sum := 0
	for _, i := range num {
		sum += i
	}
	return sum
}

func main() {
	sli := []string{"anay", "vinu", "vinayaka", "vinayaka"}

	fmt.Println(Repeater([]int{10, 100, 400, 500}))

	fmt.Printf("The rest element is %s\n", sli[0:2])
}
