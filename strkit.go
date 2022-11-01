package strkit

import "fmt"

// Ctolb CropToLenBytes crops string text to len n in characters
// regardless if the last charachter was broken or not
func Ctolb(text string, n int) string {
	return fmt.Sprintf("%s", text[0:int(n)])
}

// Ctolc CropToLenCharachters crops string text to len n in characters
func Ctolc(text string, n int) string {
	r := []rune(text)
	return fmt.Sprintf("%s", string(r[0:int(n)]))
}

// Ctolweb CropToLenWithEndBytes takes string, length and ending, and returns
// string which's byte length equals or less than n
func Ctolweb(text string, n int, ending string) string {
	tb := []byte(text)
	return fmt.Sprintf("%s%s", string(tb[0:int(n)-len(ending)]), ending)
}

// Ctolwec CropToLenWithEndCharachters takes string, length and ending, and returns
// string which's charachter amount equals to n
// Newline will be treated as one character
func Ctolwec(text string, n int, ending string) string {
	return fmt.Sprintf("%s%s", Ctolc(text, n-LenChar(ending)), ending)
}

// LenByte returns size of string in bytes
func LenByte(s string) int {
	return len(s)
}

// LenChar returns amount of characters in string
func LenChar(s string) int {
	r := []rune(s)
	return len(r)
}

// Verifiers

// Fitsc returns true if length of s is
// less than or equals n (in charachters)
//
// In other words, Fitsc returns true if s fits in n charachters
func Fitsc(s string, n int) bool {
	if LenChar(s) <= n {
		return true
	}
	return false
}

// Fitsb returns true if amount of bytes in s
// is less than or equals to n
func Fitsb(s string, n int) bool {
	if LenByte(s) <= n {
		return true
	}
	return false
}
