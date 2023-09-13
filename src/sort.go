package main

import (
	"bytes"
	"log"
	"os"
	"sort"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	log.Printf("Opening %s\n", inputFilePath)
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Printf("Error opening %s: %s\n", inputFilePath, err)
	}

	log.Printf("Reading from %s\n", inputFilePath)
	records := [][]byte{}
	for {
		rec := make([]byte, 100)
		n, err := inputFile.Read(rec)
		if err != nil {
			log.Printf("Error reading from %s: %s\n", inputFilePath, err)
		}
		if n == 0 {
			break
		}
		records = append(records, rec)
	}

	log.Printf("Closing %s\n", inputFilePath)
	inputFile.Close()

	log.Printf("Sorting %s to %s\n", inputFilePath, outputFilePath)
	sort.Slice(records, func(i, j int) bool {
		return bytes.Compare(records[i][:10], records[j][:10]) == -1
	})

	log.Printf("Creating %s\n", outputFilePath)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Printf("Error creating %s: %s\n", outputFilePath, err)
	}

	log.Printf("Writing to %s\n", outputFilePath)
	for i := 0; i < len(records); i++ {
		outputFile.Write(records[i])
	}

	log.Printf("Closing %s\n", outputFilePath)
	outputFile.Close()
}
