package main

import (
	"bufio"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func HashObject(args []string, fileName string, fileCreate bool) {

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	buffReader := bufio.NewReader(file)
	content, err := io.ReadAll(buffReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %v\n", err)
		os.Exit(1)
	}

	header := []byte(fmt.Sprintf("blob %d\x00", len(content)))
	filContent := append(header, content...)
	hashContent := sha1.Sum([]byte(filContent))
	fmt.Printf("%x\n", hashContent)
	if !fileCreate {
		return
	}

	filePath := fmt.Sprintf(".git/objects/%x", hashContent[:1])
	newFilePath := fmt.Sprintf(".git/objects/%x/%x", hashContent[:1], hashContent[1:])
	err = os.MkdirAll(filePath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating directory %v\n", err)
		os.Exit(1)
	}

	file, err = os.Create(newFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating file %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	zlibWriter := zlib.NewWriter(writer)

	_, err = zlibWriter.Write([]byte(filContent))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to file %v\n", err)
		os.Exit(1)
	}

	zlibWriter.Close()
	writer.Flush()
}
