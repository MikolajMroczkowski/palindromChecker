package main

import "fmt"

func isPalindrom(string string) bool {
	revString := reverseString(string)
	return revString == string
}

func reverseString(s string) string {
	a := []byte(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	return string(a)
}

func main() {
	var palindrom string
	fmt.Scanln(&palindrom)
	fmt.Println(isPalindrom(palindrom))
}
