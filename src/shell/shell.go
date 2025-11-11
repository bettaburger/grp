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

		// some problems with running file, clicking enter to run grp//fixed
	 
		arg := Parse(line, MAXARG)		// PARSE EACH LINE

		fmt.Println(arg.cmd, arg.args)
		/*for i, arg := range args {
    		fmt.Printf("	printed: arg[%d] = %s\n", i, arg)
		}*/

		// exit command
		if (line == "exit") {						
			fmt.Printf("	>>exiting grp shell<<\n")
			break
		}

		//Execute(line)

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