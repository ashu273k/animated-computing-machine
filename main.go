package main

import "fmt"

func frames() []string {
	return []string{"[=     ]", "[==    ]", "[===   ]", "[ ==== ]", "[  === ]", "[   == ]", "[    = ]"}
}

func main() {
	for _, frame := range frames() {
		fmt.Println(frame)
	}
}
