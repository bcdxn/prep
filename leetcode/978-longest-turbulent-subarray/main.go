package main

import "fmt"

func main() {
	nums := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	max := maxTurbulenceSize((nums))
	fmt.Println("max: ", max)
}

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	maxLen := 1
	prevSign := ""
	l := 0
	r := 1

	for r < len(arr) {
		sign := getSign(arr[r-1], arr[r])

		if sign != prevSign && sign != "=" {
			currLen := r - l + 1
			if currLen > maxLen {
				maxLen = currLen
			}
		} else {
			if sign == "=" {
				l = r
			} else {
				l = r - 1
			}
		}

		prevSign = sign
		r += 1
	}

	return maxLen
}

func getSign(a, b int) string {
	if a < b {
		return "<"
	} else if a > b {
		return ">"
	} else {
		return "="
	}
}
