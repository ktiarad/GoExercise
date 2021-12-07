package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var tempPath = filepath.Join(yourpath, "PipelineTemp")

func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(), "seconds")
}

func proceed() {
	counterTotal := 0
	counterRenamed := 0

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {

		// If there is an error, return immediately
		if err != nil {
			return err
		}

		// If it is a sub directory, return immediately
		if info.IsDir() {
			return nil
		}

		counterTotal++

		// Read file
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Sum it
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// Rename file
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("ERROR :", err.Error())

		log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
	}
}
