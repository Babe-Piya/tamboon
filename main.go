package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"Babe-Piya/tamboo/adapter/rest/payment"
	"Babe-Piya/tamboo/cipher"
	"Babe-Piya/tamboo/config"
	"Babe-Piya/tamboo/service"
)

func main() {
	conf := config.LoadConfig("config.yaml")
	paymentAPI := payment.NewOmiseAPI(conf.Omise.PublicKey, conf.Omise.SecretKey)
	songPahPaService := service.NewSongPahPaService(paymentAPI)

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, fileInfo.Size())
	fileDec, err := cipher.NewRot128Reader(file)
	if err != nil {
		fmt.Println(err)
	}

	for {
		_, errRead := fileDec.Read(buffer)
		if errRead != nil {
			if errRead == io.EOF {
				break
			}
			log.Fatalf("Error reading chunk: %v", err)
		}
	}

	reader := bytes.NewReader(buffer)
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if len(records) > 0 {
		records = records[1:]
	}

	songPahPaService.Donate(records)

}
