package main

import (
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrorLoadingFile  = errors.New("Did not found object")
	ErrorReadingFile  = errors.New("Could not read object")
	ErrorObjectConcat = errors.New("Weird error coming from reading object")
)

func ToeCatFile() error {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "usage: mygit cat-file  <object_name>\n")
		os.Exit(1)
	}

	fileName := os.Args[3]
	objectFolder := fileName[0:2]
	objectName := fileName[2:]
	filePath := ".git" + "/objects/" + objectFolder + "/" + objectName

	objectBuffer, err := os.Open(filePath)
	defer objectBuffer.Close()
	if err != nil {
		return ErrorLoadingFile
	}

	uncompressedObject, err := zlib.NewReader(objectBuffer)
	defer uncompressedObject.Close()

	if err != nil {
		return ErrorReadingFile
	}

	bytes, err := io.ReadAll(uncompressedObject)

	if err != nil {
		return ErrorObjectConcat
	}

	result := string(bytes)

	fmt.Print(result[8:])

	return nil
}
