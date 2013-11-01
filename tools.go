package main

import (
	"os"
	"log"
	)

//Ask if a file exists
//TODO: verify that is not a directory
func fileExists(path string)bool{ 
	log.Printf("Verifing that file %s exists?\n",path)
	_, err := os.Stat(path);
	return !os.IsNotExist(err)
}

//Get the content of the file at path filestring in a []byte format
func _getContentOfFile(filestring string) []byte {
	file, err := os.Open(filestring)
	if err != nil {
		log.Fatal(err)
	}
	fileinfo , err := file.Stat()
	data := make([]byte, fileinfo.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}