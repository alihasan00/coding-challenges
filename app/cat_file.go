package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CatFile(args []string) {
	filePath := args[3]
	if len(filePath) < 4 {
		fmt.Fprintf(os.Stderr, "usage: git cat-file -p <blob>\n")
		os.Exit(1)
	}
	fileName := fmt.Sprintf(".git/objects/%s/%s", args[3][:2], args[3][2:])

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	zReader, err := zlib.NewReader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %s\n", err)
		os.Exit(1)
	}
	defer zReader.Close()

	zBuffReader := bufio.NewReader(zReader)
	content, err := zBuffReader.ReadString('\x00')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %s\n", err)
		os.Exit(1)
	}
	parts := strings.SplitN(content[:len(content)-1], " ", 2)
	if len(parts) < 2 {
		fmt.Fprintf(os.Stderr, "Invalid object header\n")
		os.Exit(1)
	}
	if parts[0] != "blob" {
		fmt.Fprintf(os.Stderr, "Invalid object type: expected blob hash, got %s\n", parts[0])
		os.Exit(1)
	}
	rest := parts[1]
	parts2 := strings.SplitN(rest, "\x00", 2)
	size := parts2[0]
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %s\n", err)
		os.Exit(1)
	}

	buff := make([]byte, sizeInt)

	byteSize, err := zBuffReader.Read(buff)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %s\n", err)
		os.Exit(1)
	}
	if byteSize != sizeInt {
		fmt.Fprintf(os.Stderr, "Invalid size of blob: expected %d, got %d\n", sizeInt, byteSize)
		os.Exit(1)
	}
	fmt.Printf("%s", string(buff[:byteSize]))
}
