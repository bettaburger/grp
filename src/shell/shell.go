package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pborman/getopt/v2"
)

var (
	debug int
)

//accepts up to 100 arguments
const MAXARG int = 100

/* BASH SHELL */
func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for ((true)) {
		fmt.Printf("grp$ ")

		if !scanner.Scan() {
			break
    	}

		line := strings.TrimSpace(scanner.Text())	// strips newline '\n'
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
    	} 
		// PARSE EACH LINE
		arg := Parse(line, MAXARG)		//fmt.Printf("	%s	%s	\n",arg.cmd, arg.args)
		
		// calling execute
		Execute(arg.cmd)
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