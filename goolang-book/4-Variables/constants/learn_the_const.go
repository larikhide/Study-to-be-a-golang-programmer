package main

import "fmt"

func main() {
	const x string = "Hello Wolrd" /* значения констант не могут быть изменены после их иницаиализации. Константы - хороший способ использовать определенные значения в программе без необходимости писать их каждый раз */
	fmt.Println(x)
}