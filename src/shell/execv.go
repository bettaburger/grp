package shell

import (
	"fmt"
)

/*str comparison function*/
func strCmp(a string, b string) int {
	switch {
		// if equal return 0
	case a == b:
		return 0
		// otherwise return -1
	default:
		return -1	
	}
}

/*EXECUTE FUNCTION
	if statements for commands
*/
func Execute(command string) {
	for ((true)) {
		// exit command
		if strCmp(command, "exit") == 0 {
			break
		// ls command
		} else if strCmp(command, "ls") == 0 {
			// call ls command from CMDS here
			fmt.Printf("	run ls here..\n")
			break
		}
	}
}