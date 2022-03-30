package main

import (
	"bufio"
	"os"
)

// Take file where write as parameter and string to write
// And write in the file without overwriting what is written in before
func writeInFile(input string, fPtr *os.File) {
	fPtr.Sync()
	w := bufio.NewWriter(fPtr)
	_, err := w.WriteString(input + "\n")
	debugmsg("BUFFER= " + input)
	warningerr(err)
	// flush the buffer to free memory
	w.Flush()
}


func openFileToWrite(ofname string) *os.File {
	// Open file or create it if not already exist
	fd, err := os.OpenFile(ofname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	fatalerr(err)
	debugmsg("file open")
	return fd
}

func closeFileToWrite(f *os.File) {
	// close file before exit function
	f.Close()
	debugmsg("file close")
}

// Copyright (C) 2020 Guillaume GRONNIER. All rights reserved
