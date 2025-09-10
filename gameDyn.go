package cyoa

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Functions
// Defining functions for the dynamics of the game

func Read_story_from_json() []byte {
	// Default File json file to extract story
	default_file := "gopher.json"

	// In case the user wants to input another story
	story_file := flag.String("file", default_file, "JSON file name which contains the story of the game/book")
	// Loading the inputed value or the default file name
	flag.Parse()

	// Openning the file
	stories_folder_path := "Stories_JSON"
	path_to_file := filepath.Join(stories_folder_path, *story_file)
	data, err := os.ReadFile(path_to_file)
	if err != nil {
		panic(err)
	}

	return data
}

// Parse JSON struct into a map of Arc structure
func ParseJson(json_byte []byte, story *Story) {
	err := json.Unmarshal(json_byte, &story)
	if err != nil {
		panic(err)
	}
}

// Despliega solo el texto del arco (capítulo) de un elemento JSON
func ReadChapterText(paragraphs []string) {
	for _, paragraph := range paragraphs {
		fmt.Println(paragraph)
		fmt.Println("")
	}
}

// Display and ask user to choose on of the chapter options (whic is another chapter to move on)
func ChoosingOption(options []Option) string {
	fmt.Println("---------Available Options:---------\n")
	for n, option := range options {
		fmt.Printf("Option #%d: %s\n", n, option.Text)
	}

	fmt.Println("\n Type the number of the selected option:")

	// Receiving input
	var selected_option int8
	fmt.Print("-->")
	// TBD I can put restrictions of the tyoe pf values to be inputes e.g. not alfa rokman skull
	fmt.Scanf("%d\n", &selected_option)

	// Returns the index of the selected option
	selected_chapter := options[selected_option].Arc
	return selected_chapter
}

// Loop to display other chapter after the initial one
func InitGame(story Story, current_chapter *string) {
	for game_continues := true; game_continues; {

		fmt.Printf("\n---------Chapter: %s---------\n", story[*current_chapter].Title)
		ReadChapterText(story[*current_chapter].Story)
		if len(story[*current_chapter].Options) > 0 {
			*current_chapter = ChoosingOption(story[*current_chapter].Options)
		} else {
			fmt.Println("THE GAME ENDS HERE.\nThanks for playing!")
			game_continues = false
		}
	}
}

// Structures
// Defining structures to store data from JSON (remember flags for decoding)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Arc
