package main

import (
	"cyoa"
	"fmt"
)

// 1. Function to initialize the program
func main() {


	// Preparing Game interface

	// 1. Load the story from the json file
	story_byte := cyoa.Read_story_from_json()

	// 2. Parse story to defined objects
	var story cyoa.Story
	cyoa.ParseJson(story_byte, &story)

	// Initializing game

	// 3. Display instructions and and the chapter title and story text
	fmt.Println("___________________________________________________")
	fmt.Println("\nChoose your own adventure, the interactive book in which you decide how the story ends\n")
	fmt.Println("After each chapter select one of the available options\n")
	fmt.Println("___________________________________________________\n")

	current_chapter := "intro"

	// 4. Display options and let user select the next chapter to interact with
	cyoa.InitGame(story, &current_chapter)

}
