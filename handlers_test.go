package cyoa

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Testread_story_from_json(t *testing.T) {
	// Default File json file to extract story
	default_file := "gopher.json"

	// In case the user wants to input another story
	story_file := flag.String("json", default_file, "JSON file name which contains the story of the game/book")
	// Loading the inputed value or the default file name
	flag.Parse()

	// Openning the file
	stories_folder_path := "Stories_JSON"
	path_to_file := filepath.Join(stories_folder_path, *story_file)
	data, err := os.ReadFile(path_to_file)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	return 
	// checking result
	// if result["intro"].Title != "The Little Blue Gopher" {
	// 	t.Errorf("Expected title 'The Little Blue Gopher', got '%s'", result["intro"].Title)
	// }
}
