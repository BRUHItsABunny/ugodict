package main

import (
	"fmt"
	"strconv"
	"ugodict"
)

func main() {
	// Init client
	client := ugodict.GetClient()

	// Get definition
	definition, err := client.DefineById("6267709")

	// Check error
	if err == nil {
		// Print data
		fmt.Println("ID: " + strconv.Itoa(definition.DefId))
		fmt.Println("Author: " + definition.Author)
		fmt.Println("Word: " + definition.Word)
		fmt.Println("Definition: " + definition.Definition)
		fmt.Println("Example: " + definition.Example)
		fmt.Println("Thumbs up: " + strconv.Itoa(definition.ThumbsUp))
		fmt.Println("Thumbs down: " + strconv.Itoa(definition.ThumbsDown))
		fmt.Println("Permalink: " + definition.PermaLink)
		fmt.Println("Written on: " + definition.WrittenOn)
	} else {
		// Print error
		fmt.Println(err)
	}
}
