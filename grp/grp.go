package grp

import (
	"fmt"
	//"os"
	"github.com/pborman/getopt/v2"
)

var (
	debug int
)

func Grp() {
	// flag commands
	dFlag := getopt.Bool('d', "enable debug")	// -d debug
	// parse the arguments
	getopt.ParseV2()
	
	switch {
	case *dFlag:
		debug = 1
		fmt.Println("Debug enabled")	
	default:
		getopt.Usage()	
	}
	
	// test case for debug
	for debug == 1 {
		fmt.Println("hello this debug is 1")
		break;
	}


}