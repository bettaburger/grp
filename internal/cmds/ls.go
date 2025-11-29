package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"grp/internal/env"
)

/* LISTS CONTENTS IN A FOLDER/DIRECTORY */
func List(target string) {
	currPath := PWD()
	
	homeVar := env.GetEnv("HOME")
	// HOME VARIABLE
	home := strings.Split(homeVar, "=")

	// HOME VALUE
	path := home[1]

	// MANAGE TARGET
	switch target {
		// LISTS UNDER HOME, doesn't matter which directory you are in.
		// tested already with linux
		case "~", "$HOME":
			target = path

		case "..":	
			target = filepath.Dir(currPath)

		// LISTS UNDER CURRENT DIRECTORY
		case ".", "":
			target = currPath
		
		// LISTS A SPECIFIC DIRECTORY/PATH, doesn't matter which directory you are in
		default:
		// If target is a relative path, join it to cwd
        if !filepath.IsAbs(target) {
            target = filepath.Join(currPath, target)
        }
    }

    // PATH OR DIRECTORY NOT FOUND, print error. "no such file or directory"
    files, err := os.ReadDir(target)
    if err != nil {
        fmt.Println("no such file or directory")
        return
    }

	// PRINT THE FILES
    for _, file := range files {
        fmt.Println(file.Name())
    }
}