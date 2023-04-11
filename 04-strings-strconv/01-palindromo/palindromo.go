package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	arrayInverse := make([]string, 0)
	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInverse = append(arrayInverse, arrayCadena[i])
	}
	return strings.Join(arrayInverse, "")
}
func isPalindrome(word string) bool {
	newWord := strings.ToLower(word)
	newWord = strings.ReplaceAll(newWord, " ", "")
	inverse := reverse(newWord)

	return newWord == inverse
}

func main() {
	isPalindrome := isPalindrome("Anita lava la tina")
	if isPalindrome {
		fmt.Println("SI ES PALINDROMO")
	} else {
		fmt.Println("No ES PALINDROMO")
	}
}
