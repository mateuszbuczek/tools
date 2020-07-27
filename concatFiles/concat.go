package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Format: <dir_name> <output_filename>")
	}

	folderName, outputFileName := os.Args[1], os.Args[2]

	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	result, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	log.Printf("Found %d files", len(files))
	writeFilesToFile(files, folderName, result)
}

func writeFilesToFile(files []os.FileInfo, folderName string, result *os.File) {
	for _, f := range files {
		file, err := os.Open(folderName + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		writeFileToFile(file, result)
	}
}

func writeFileToFile(file *os.File, result *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		_, err := result.Write(append(scanner.Bytes(), '\n'))
		if err != nil {
			log.Fatal(err)
		}
	}
}
