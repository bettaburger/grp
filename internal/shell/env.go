package shell

import (
	"os"
	"fmt"
	"strings"
	"grp/pkg/compare"
)

var (
	env []string = os.Environ()		// list of environments
)

/*ENVP FOR PATHING*/
type Envp struct {
	Variable 	string		// key
	Value 	[]string	// value
}
/*ENVIRONMENT FUNCTION
	temporary
	move this to execv.go, initialize env 
*/
func Env() {
	for i, e := range env {
		fmt.Println(i, "| ", e)
	}
}

/* RETRIEVES ENVIRONMENT */
func GetEnv(key string) {
	var found int
	for _, e := range env {
		variables := strings.Split(e, "=")
		// if key is found, print the environment
		if compare.StrCmp(key, variables[0]) == 0 {
			found = 1
			fmt.Printf("%s\n", e) 
		} 
	}
	if found == 0 {
		fmt.Printf("%s environment not found\n", key)
	}
}

/*UPDATES ENVIRONMENT
	takes a key and changes the variables 
*/
func setEnv(key string, variable string) {
}
