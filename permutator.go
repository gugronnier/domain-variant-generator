package main

import (
	"flag"
)

func main() {
	var inDom, ofPtr string
	var debugPtr, fullverbosePtr bool
	
	flag.StringVar(&inDom, "d", "exemple.com", "Domain")
	flag.StringVar(&ofPtr, "of", "output.csv", "Output File wanted")
	flag.BoolVar(&debugPtr, "v", false, "enable debug messages level")
	flag.BoolVar(&fullverbosePtr, "vv", false, "enable debug, warning and info messages level")
	flag.Parse()

	// Set verbosity level
	if debugPtr {
		verbosity = 1
	}
	if fullverbosePtr {
		verbosity = 2
	}
	if ofPtr != defaultoutputfile {
		outputfile = openFileToWrite(ofPtr)
	}else{
		outputfile = openFileToWrite(defaultoutputfile)
	}
	
	// add title to csv
	textToWrite := "Domain,Is Register,Registrar"
	writeInFile(textToWrite, outputfile)
	
	checkargs(flag.NFlag(),inDom)
	
	// End
	if outputinfile{
		closeFileToWrite(outputfile)
	}
}


