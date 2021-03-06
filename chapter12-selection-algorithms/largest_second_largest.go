// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty without even the implied warranty of
//

package main

import (
	"fmt"
	"math"
)

func findLargestAndSecondLargest(A []int) (largest, second int) {
	largest, second = math.MinInt32, math.MinInt32
	for _, v := range A {
		if v > largest {
			second = largest
			largest = v
		} else if v > second {
			second = v
		}
	}
	return largest, second
}

func main() {
	largest, second := findLargestAndSecondLargest([]int{11, -4, 7, 8, -10})
	fmt.Println("largest: ", largest, "second: ", second)
}
