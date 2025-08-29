package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"Babe-Piya/tamboo/cipher"
)

func main() {
	fmt.Println("Hello World")
	file, err := os.Open("data/fng.1000.csv.rot128")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("File Size: ", fileInfo.Size())

	buffer := make([]byte, fileInfo.Size())
	fileDec, err := cipher.NewRot128Reader(file)
	if err != nil {
		fmt.Println(err)
	}

	for {
		n, errRead := fileDec.Read(buffer)
		if errRead != nil {
			if errRead == io.EOF {
				break
			}
			log.Fatalf("Error reading chunk: %v", err)
		}
		fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
	}

	fmt.Printf("end Read %d bytes: %s\n", len(buffer), string(buffer))

}
