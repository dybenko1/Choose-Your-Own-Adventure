package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// 1. Function to initialize the program
func main() {

	// 0. Initializing the port for the game to be launc
	port := flag.Int("flag", 3000, "Port to be launch CYOA")

	// Preparing Game interface

	// 1. Load the story from the json file
	story_byte := cyoa.Read_story_from_json()

	// 2. Parse story to defined objects
	var story cyoa.Story
	cyoa.ParseJson(story_byte, &story)

	// Initializing game

	// 3. Launching game in a webapp
	h := cyoa.NewHandler(story)
	fmt.Printf("Launching on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

	// Code to play CYOA in the terminal
	// // 3. Display instructions and and the chapter title and story text
	// fmt.Println("___________________________________________________")
	// fmt.Println("\nChoose your own adventure, the interactive book in which you decide how the story ends\n")
	// fmt.Println("After each chapter select one of the available options\n")
	// fmt.Println("___________________________________________________\n")

	// current_chapter := "intro"

	// // 4. Display options and let user select the next chapter to interact with
	// cyoa.InitGame(story, &current_chapter)

}
