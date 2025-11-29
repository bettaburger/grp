package cmds

import (
	"fmt"
	"grp/internal/env"
	"os"
	"path/filepath"
	"strings"
)

/*PRINTS THE CURRENT PATH*/
func PWD() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("cannot get current PATH")
		
	}
	return dir
}
/*CHANGE DIRECTORY*/
// cd 			==> cd home
// cd ~ 		==> cd home
// cd $HOME 	==> cd home

// cd .. 		==> cd back

func ChangeDir(target string) {
	currPath := PWD()
	
	homeVar := env.GetEnv("HOME")
	// HOME VARIABLE
	home := strings.Split(homeVar, "=")

	// HOME VALUE
	path := home[1]

	// MANAGE TARGET
	switch target {
	case "", "~", "$HOME":
		target = path
		// parent directory
	case "..":
		target = filepath.Dir(currPath)
	default:
		// if target starts with ~, it is a home path
		if strings.HasPrefix(target, "~") {
			target = strings.Replace(target, "~", path, 1)
		}
		// not a full path
		if !filepath.IsAbs(target) {
			target = filepath.Join(currPath, target)
		}
	}

	// Changes directory
	if err := os.Chdir(target); err != nil {
		fmt.Printf("no such file or directory\n")
		return
	}

	// update OLDPWD AND PWD
	newDir := PWD()
	env.SetEnv("OLDPWD", currPath)
	env.SetEnv("PWD", newDir)
}