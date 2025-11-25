package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func readBytes(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes in the file: %d\n", len(content))
}

func readFileLines(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count += 1
	}
	log.Printf("Number of lines in the file: %d\n", count)
}

func readNumberOfWords(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		words := strings.Fields(lineText)
		count += len(words)
	}
	log.Printf("Number of words in the file: %d\n", count)
}

func readNumberOfCharacters(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewReader(file)
	count := 0
	for {
		_, _, err := scanner.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		count += 1

	}
	log.Printf("Number of Characters in the file: %d\n", count)
}
