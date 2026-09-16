package main

import "fmt"

func checkSign(n int) string {
	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	}
	return "Zero"
}

func main() {
	fmt.Println(checkSign(15))
	fmt.Println(checkSign(-7))
	fmt.Println(checkSign(0))
}
