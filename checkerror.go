package main

import(
	"os"
)

// print fatal error if there is and exit
func fatalerr(e error) {
	if e != nil {
		ErrorColor.Println("[FATAL] " + e.Error())
		ErrorColor.DisableColor()
		os.Exit(1)
	}
}

// print error as a warning if there is and return
func warningerr(e error) bool {
	if e != nil {
		if verbosity == 2 {
			WarningColor.Println("[WARNING] " + e.Error() + "\n")
			WarningColor.DisableColor()
		}
		return true
	}
	return false
}

// check verbosity level and print debug message if verbose mode is enable
func debugmsg(msg string) {
	if verbosity != 0 {
		DebugColor.Println("[DEBUG] " + msg)
		DebugColor.DisableColor()
	}
}

// manually raise fatal error
func errormsg(s string) {
	ErrorColor.Println("[ERROR] " + s)
	ErrorColor.DisableColor()
}

// Copyright (C) 2020 Guillaume GRONNIER. All rights reserved
