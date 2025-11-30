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
	fmt.Printf("grp help			show this message and then exit\n")
	fmt.Printf("exit				stop the process")
	fmt.Print("\n")
	fmt.Print("\n")
	// Print basic commmands 
	color.Set(color.FgBlue)
	fmt.Println(">> Basic commands:")
	color.Unset()
	// list of basic commands
	fmt.Printf("ls\ncd\nenv\n")
	/* spacing */
	fmt.Print("\n")
}