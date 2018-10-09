// Fizzbuzz for Zenvia, based on http://codingdojo.org/kata/FizzBuzz/
// Fizzbuzz is a program that prints interger numbers from 1 to n. But for multiples of three,
// print “Fizz” instead of the number and for the multiples of five print “Buzz”.
// For numbers which are multiples of both three and five print “FizzBuzz”.
// Fizzbuzztwo does the same, but print "Fizz" and "Buzz" when the number has 3 and/or 5 in it too.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter an integer number: ")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(i int) string {
	var result string
	var istr = strconv.Itoa(i)

	if i%3 == 0 {
		result = "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}
	if len(result) == 0 {
		result = istr
	}

	return result
}

func fizzbuzztwo(i int) string {
	var result string
	var istr = strconv.Itoa(i)

	if i%3 == 0 || strings.Contains(istr, "3") {
		result = "Fizz"
	}
	if i%5 == 0 || strings.Contains(istr, "5") {
		result += "Buzz"
	}
	if len(result) == 0 {
		result = istr
	}

	return result
}
