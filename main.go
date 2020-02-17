package main

import "fmt"

func main() {

	fmt.Println("Favs:", favs())
}

func favs() []string {
	return []string{"favone"}
}
