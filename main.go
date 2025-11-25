package main

import (
	"flag"
	"log"
)

func main() {
	readByteFileName := flag.String("c", "", "print the number of bytes")
	numberOfLines := flag.String("l", "", "print the number of lines in file")
	numberOfWords := flag.String("w", "", "print the number of words in file")
	numberOfCharacters := flag.String("m", "", "print the number of characters in file")

	flag.Parse()
	switch {
	case *readByteFileName != "":
		readBytes(string(*readByteFileName))
	case *numberOfLines != "":
		readFileLines(string(*numberOfLines))
	case *numberOfWords != "":
		readNumberOfWords(*numberOfWords)
	case *numberOfCharacters != "":
		readNumberOfCharacters(*numberOfCharacters)
	default:
		log.Println("Usage: wc [-c | -l | -w | -m] <file>")
	}

}
