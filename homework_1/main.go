package main

import (
	"fmt"
)

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
	Pink  = 13
	White = 255
)

type sandglassMapa map[string]int

func sandglass(args ...sandglassMapa) {
	defaultMapa := sandglassMapa{"size": 15, "color": White, "char": 'X'}
	for _, arg := range args {
		for key, value := range arg {
			defaultMapa[key] = value
		}
	}

	size := defaultMapa["size"]
	color := defaultMapa["color"]
	char := defaultMapa["char"]

	fmt.Printf("\033[38;5;%dm", color) //set color char
	for i := 0; i < size/2; i++ {      //print up part
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}
		for k := i; k < size-i; k++ {
			if i == 0 {
				fmt.Printf("%c", char)
			} else if k == i || k == size-i-1 {
				fmt.Printf("%c", char)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
	for i := size / 2; i >= 0; i-- { //print down part
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}
		for k := i; k < size-i; k++ {
			if i == 0 {
				fmt.Printf("%c", char)
			} else if k == i || k == size-i-1 {
				fmt.Printf("%c", char)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	sandglass()
	sandglass(sandglassMapa{"color": Cyan, "size": 7, "char": '$'})
}
