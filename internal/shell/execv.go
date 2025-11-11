package shell

import (
	"fmt"
	"grp/pkg/compare"
)

/*EXECUTE FUNCTION
	if statements for commands
*/
func Execute(command string) {
	for ((true)) {
		// exit command
		if compare.StrCmp(command, "exit") == 0 {
			break
		// ls command
		} else if compare.StrCmp(command, "ls") == 0 {
			// call ls command from CMDS here
			fmt.Printf("	run ls here..\n")
			break
		} else {
			fmt.Print("command not found try --help or -h\n")
			break
		}
	}
}