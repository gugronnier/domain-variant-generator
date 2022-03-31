package main

import (
	"regexp"
)

func checkargs (i int, inDom string) {
	// check number of 
	if i > 3 {
		errormsg("Too much argument")
		runhelp()
	}
	if i < 2 {
		errormsg("Need more argument")
		runhelp()
	}
	
	if isDomain(inDom) {
		permute_string(inDom)
	} else {
		errormsg("Invalid Arguments")
		runhelp()
	}
}

func isDomain (input string) bool {
	// domain regex initialisation
	domain := regexp.MustCompile("^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if domain.MatchString(input) {
		return true
	}else{
		return false
	}
}