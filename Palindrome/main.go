package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}
func main() {
	x := 121
	fmt.Println(isPalindrome(x))
	y := -121
	fmt.Println(isPalindrome(y))
	z := 10
	fmt.Println(isPalindrome(z))
}
