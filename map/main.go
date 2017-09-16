package main

import "fmt"

func main() {
	colors := map[string]string{
		"white": "#ffffff",
		"red":   "#ff0000",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)
	}
}
