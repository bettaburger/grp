package debug

import (
	"fmt"

	"github.com/fatih/color"
)

/*LISTS ALL grp COMMANDS*/
func HelperList() {
	/* spacing */ 
	fmt.Print("\n")

	// Examples of commandline use
	color.Set(color.FgGreen)
	fmt.Println(">> Example usage:")
	color.Unset()
	// lists of examples on how to use commands
	fmt.Println("...")
	fmt.Print("\n")
	fmt.Print("\n")
	// Print basic commmands 
	color.Set(color.FgBlue)
	fmt.Println(">> Basic commands:")
	color.Unset()
	// list of basic commands
	fmt.Println("ls")
	fmt.Println("echo")

	/* spacing */
	fmt.Print("\n")
}