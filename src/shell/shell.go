package shell

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/pborman/getopt/v2"
)

var (
	debug int
)

/* BASH SHELL */
func Run() {
	// scan for user input
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		fmt.Printf("grp$ ")
		line := strings.TrimSpace(scanner.Text())	// strips newline '\n'	

		if (line == "exit") {						// Make exit call, exit
			fmt.Printf("exiting grp shell\n")
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
    	}
		// some problems with running file, clicking enter to run grp

		// echoes line back
		fmt.Printf("%s", line) 
	}
	
	

	/* DEBUGGER 
		run: $ -d
	*/
	dFlag := getopt.Bool('d', "enable debug")	// -d debug
	// parse the arguments
	getopt.ParseV2()
	
	switch {
	case *dFlag:
		debug = 1
		fmt.Println("Debug enabled")	
	default:
		getopt.Usage()	
	}
	
	// test case for debug
	for debug == 1 {
		fmt.Println("hello this debug is 1")
		break;
	}




}