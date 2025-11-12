package shell

import (
	"fmt"

	"grp/pkg/compare"
	"grp/pkg/debug"	
)

/*EXECUTE FUNCTION
	if statements for commands
*/
func Execute(command string, line string) {
	for {
		// ls command
		if compare.StrCmp(command, "ls") == 0 {
			// call ls command from CMDS here
			fmt.Printf("	run ls here..\n")
			break
		// grp help command
		} else if compare.StrCmp(line, "grp help") == 0 {
			debug.HelperList()
			break
		// echo command
		} else if compare.StrCmp(command, "echo") == 0 {
			fmt.Printf("	run echo here..\n")
			break
		// env command
		} else if compare.StrCmp(command, "env") == 0{
			Environment()
			break
		// unidentified command
		} else {
			fmt.Print("command not found try 'grp help'\n")		
			break
		}
	}
}