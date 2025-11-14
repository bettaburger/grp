package env

import (
	"fmt"
	"os"
	"strings"
	"slices"
	"grp/pkg/compare"

)

var (
	env []string = os.Environ()		// the list of environments
)

/*ENVP FOR PATHING*/
type Envp struct {
	Variable 	string		// key
	Value 	[]string	// value
}
/*ENVIRONMENT FUNCTION*/
func Env() {
	for i, e := range env {
		fmt.Println(i, "| ", e)
	}
}

/* RETRIEVES ENVIRONMENT */
func GetEnv(key string) string {
	var found int
	var variable string
	for _, e := range env {
		variables := strings.Split(e, "=")
		// if key is found, print the environment
		if compare.StrCmp(key, variables[0]) == 0 {
			found = 1
			variable = e 
			//fmt.Printf("%s\n", e)	
		} 
	}
	if found == 0 {
		fmt.Printf("%s environment not found\n", key)
	}
	return variable
}

/*UPDATES ENVIRONMENT
	takes a key and changes the value
*/
func SetEnv(key string, value string) {
	var found int 
	for i, e := range env {
		variables := strings.Split(e, "=")
		// if key is found, print the environment
		if compare.StrCmp(key, variables[0]) == 0 {
			found = 1
			// change the old environment variable to the new variable
			env[i] = fmt.Sprintf("%s=%s", key, value)
			break
		} 
	}
	// if key is not an environment, make a new one
	if found == 0 {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
}
/*REMOVES ENVIRONMENT
	make sure to shift all the variables to the right. 
*/
func UnsetEnv(key string) {
	var found int 
	for i, e := range env {
		// key = variables
		variables := strings.Split(e, "=")
		// if key is found, print the environment
		if compare.StrCmp(key, variables[0]) == 0 {
			found = 1
			// remove the environment
			env = slices.Delete(env, i, i+1)
			break
		} 
	}
	// key does not exist
	if found == 0 {
		fmt.Printf("%s environment not found\n", key)
	}
	
}
