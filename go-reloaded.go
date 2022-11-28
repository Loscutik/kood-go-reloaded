package main

import (
	//"fmt"
	"bufio"
	"log"
	"os"

	"01.kood.tech/git/obudarah/go-reloaded/modifytext"
)

func main() {
	// needs 2 arguments
	if len(os.Args) != 3 {
		log.Fatalln("Wrong numbers of arguments. Needs 2 files")
	}

	fileIn, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error when openning the file: %s", err)
	}
	defer fileIn.Close()

	fileOut, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("Error when creating the file %s", err)
	}
	defer fileOut.Close()

	// read fileIn line by line
	fileScanner := bufio.NewScanner(fileIn)
	for fileScanner.Scan() { // return false when the scan stops, either by reaching the end of the input or an error
		str, err := modifytext.ModifyString(fileScanner.Text())
		if err != nil {
			log.Fatalf("Error while solving issues. The last is %s", err)
		}

		_, err = fileOut.Write([]byte(str + "\n"))
		if err != nil {
			log.Fatalln("Error while writing to the result file")
		}
	}

	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading the file: %s", err)
	}
	
	// delete an empty line in fileOut 
	stat, err := fileOut.Stat()
	if err != nil {
		log.Fatalf("Error while reciving stat of the file: %s", err)
	}
	if err = fileOut.Truncate(stat.Size() - 1); err != nil {
		log.Fatalf("Error while truncate an empty linr from the file: %s", err)
	}
}
