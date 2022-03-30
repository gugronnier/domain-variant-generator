package main

import (
	"github.com/fatih/color"
	"os"
)

// ## GLOBAL VARIABLES ##

// verbosemode
//	 0 = no debug output (default)
//	 1 = only debug messages and critical error
//	 2 = debug messages, warnings and critical error
var verbosity = 0
var defaultoutputfile = "output.csv"
var outputinfile = true
var outputfile *os.File

// ## CONSTANT VARIABLES ##

// color output variable
var DebugColor      = color.New(color.FgBlue) // blue
var WarningColor    = color.New(color.FgYellow) // yellow
var ErrorColor      = color.New(color.FgRed) // red
var OutputColor     = color.New(color.FgWhite) // white
var UsageColor      = color.New(color.FgGreen) // green
