package cmds

import (
	"fmt"
	"os"
)

/* take the current PWD, and list out the contents. */
func List(target string) {

	switch target {
		// LISTS UNDER HOME, doesn't matter which directory you are in.
		// tested already with linux
		case "~", "$HOME":

		// LISTS UNDER CURRENT DIRECTORY
		case ".", "":

		default:
			// LISTS A SPECIFIC DIRECTORY/PATH, doesn't matter which directory you are in
			// PATH OR DIRECTORY NOT FOUND, print error. "no such file or directory"
	}
		
	
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("error")
	}

	// lists out files/directories
	for _, file := range files {
		fmt.Println(file.Name())
	}
}