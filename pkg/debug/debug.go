package debug 

import (
	"fmt"
	"github.com/pborman/getopt/v2"
)

var (
	debug int
)

func Debug(){
	// flags
	dFlag := getopt.Bool('d', "enable debug")	// -d debug
	hFlag := getopt.Bool('h', "helper")			// -h help
	// parse the arguments
	getopt.ParseV2()
	
	switch {
	case *dFlag:
		debug = 1
		fmt.Println("Debug enabled")
	case *hFlag:
		// call helper function
		HelperList()
	default:
		getopt.Usage()	
	}
	
	// test case for debug
	for debug == 1 {
		fmt.Println("hello this debug is 1")
		break;
	}
}