package main

import (
	"fmt"
	//"os"
	"github.com/pborman/getopt/v2"
)

var debug int

func main() {
	/**
	Flag commands
	Parse command arguments
	Switch statement
	*/
	dFlag := getopt.Bool('d', "enable debug")	// -d debug
	getopt.ParseV2()
	
	switch {
	case *dFlag:
		debug = 1
		fmt.Println("debugging")	
	default:
		getopt.Usage()	
	}	
}