package shell

import (
	"os"
	"fmt"
)

/*ENVP FOR PATHING*/
type Envp struct {
	Name 	string
	Var 	[]string 
}
/*ENVIRONMENT FUNCTION
	Environment variables
	Can display the name and variable(s)
*/
func Environment() {
	env := os.Environ() // returns []string of "key=value"
	for _, e := range env {
		fmt.Println(e)
	}
	fmt.Printf("\n%d env variables\n", len(env))
}
