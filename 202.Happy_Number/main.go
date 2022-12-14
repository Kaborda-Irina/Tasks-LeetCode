package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}
func isHappy(n int) bool {
	a := getSum(n)
	if a == 1 {
		return true
	} else if a == 4 {
		return false
	} else {
		return isHappy(a)
	}
}
func getSum(n int) int {
	if n/10 == 0 {
		return n * n
	}
	return (n%10)*(n%10) + getSum(n/10)
}
