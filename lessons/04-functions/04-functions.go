package main

import "fmt"

func main() {
    text1 := ""
	text2 := ""

	text1 = "Hello\n"
	text2 = "World!\n"	

	outputText(text1, text2)
}

func outputText(textString1 string, textString2 string) {
	fmt.Print(textString1)
	fmt.Print(textString2)
}