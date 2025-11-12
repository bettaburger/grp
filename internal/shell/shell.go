package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//accepts up to 100 arguments
const MAXARG int = 100

/* BASH SHELL */
func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("grp$ ")

		if !scanner.Scan() {
			break
    	}

		line := strings.TrimSpace(scanner.Text())	// strips newline '\n'
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
    	} 
		// PARSE EACH LINE into arg
		arg := Parse(line, MAXARG)		//fmt.Printf("	%s	%s	\n",arg.cmd, arg.args)
		// exit command
		if arg.cmd == "exit" {
			break
		}
		// calling execute
		Execute(arg.cmd, line)
	}
}