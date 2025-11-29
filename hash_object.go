package main

import (
	"bufio"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func HashObject(args []string, fileCreate bool) {
	fileName := args[3]
	if fileName == "" {
		fmt.Fprintf(os.Stderr, "usage: git hash-object -w <file-name>\n")
		os.Exit(1)
	}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "something went wrong %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	buffReader := bufio.NewReader(file)
	content, err := io.ReadAll(buffReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "something went wrong %v\n", err)
		os.Exit(1)
	}

	filContent := fmt.Sprintf("blob %d\x00%s", len(content), content)
	hashContent := sha1.Sum([]byte(filContent))
	fmt.Printf("%x\n", hashContent)
	if !fileCreate {
		return
	}

	filePath := fmt.Sprintf(".git/objects/%x", hashContent[:1])
	newFilePath := fmt.Sprintf(".git/objects/%x/%x", hashContent[:1], hashContent[1:])
	os.MkdirAll(filePath, 0755)

	file, err = os.Create(newFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "something went wrong %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	zlibWriter := zlib.NewWriter(writer)

	_, err = zlibWriter.Write([]byte(filContent))
	if err != nil {
		fmt.Fprintf(os.Stderr, "something went wrong %v\n", err)
		os.Exit(1)
	}

	zlibWriter.Close()
	writer.Flush()
}
