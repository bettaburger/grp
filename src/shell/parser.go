package shell

import (
	"strings"
	"fmt"
	"os"
)

type Command struct {
	cmd 	[]string
	args 	[]string
}

/*	PARSE FUNCTION:
		Takes in a string and returns the tokens.
	EXAMPLE:
	input: ls filename
	lexing => split each word into tokens:
	1) [ls] 
	2) [filename]
	output: tokens = []string{"ls", "filename"}
*/
func Parse(line string, maxarg int) *Command {
	line = strings.TrimSpace(line)
	// lexing happens 
	tokens := strings.Fields(line)		// splits by spaces

	// Check if the length of tokens exceed 100. 
	if len(tokens) > maxarg {
		fmt.Print("Too many arguments")
		os.Exit(1)
	}
	// initialize Commands
	var cmdArg Command
	cmdArg.cmd = tokens[:1]
	cmdArg.args = tokens[1:]

	return &cmdArg
}

