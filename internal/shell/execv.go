package shell

import (
	"fmt"
	"strings"
	"grp/pkg/compare"
	"grp/pkg/debug"	
	"grp/internal/env"
	"grp/internal/cmds"
)

var (
	key string
	value string
)

/*in future, move similar commands to its own folder? reduces flow of if statements? :p
	maybe keywords such as a command's first letter or something similar. 
*/

/*EXECUTE FUNCTION
	if statements for commands
*/
func Execute(command string, line string) {
	args := strings.Fields(line)
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

/*PATHING COMMANDS*/
		// pwd command
		} else if compare.StrCmp(command, "pwd") == 0 {
			fmt.Printf("%s\n", cmds.PWD())
			break

/*CD*/
		} else if compare.StrCmp(command, "cd") == 0 {
			// check if there is "". 
			//home := env.GetEnv(args[1])
			// otherwise run cd [directory]	
			target := "~" // default: go to home
			if len(args) > 1 {
				target = args[1]
			}
			cmds.ChangeDir(target)
			break
		} else if compare.StrCmp(command, "cd ~") == 0 {
			cmds.ChangeDir("~")
		
		} else if compare.StrCmp(command, "cd $HOME") == 0 {
			cmds.ChangeDir("$HOME")

/*ENVIRONMENT COMMANDS*/
		// env command
		} else if compare.StrCmp(command, "env") == 0{
			// prints all environment variables
			env.Env()
			break
		// getenv command
		} else if compare.StrCmp(command, "getenv") == 0 {
			// gets the environment
			e := env.GetEnv(args[1])
			fmt.Printf("%s\n", e)
			break
		// setenv command
		} else if compare.StrCmp(command, "setenv") == 0 {
			key = args[1]
			value = strings.Join(args[2:], " ")
			env.SetEnv(key, value)
			break
		// unsetenv command
		} else if compare.StrCmp(command, "unsetenv") == 0 {
			key = args[1]
			env.UnsetEnv(key)
			//fmt.Println(len(env)) //==> checks to see if the env array is getting decremented by every unset
			break

/*IF IT IS NOT A COMMAND*/
		} else {
			fmt.Print("command not found try 'grp help'\n")		
			break
		}
	}
}