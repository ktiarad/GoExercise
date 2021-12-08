package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(yourpath, "PipelineTemp")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("Start")
	start := time.Now()

	generateFiles2()

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(), "seconds")
}

func randomString2(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles2() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	// Pipeline 1: Job distributin
	chanFileIndex := generateFileIndexes()

	// Pipeline 2: The main logic (creating files)
	createFilesWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFilesWorker)

	// Tranc and print output
	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("Error creating file %s. Stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("File-%d.txt", i),
			}
		}
		close(chanOut)
	}()
	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// Wait group to control the workers
	wg := new(sync.WaitGroup)

	// Allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		// Dispatch N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				// Listen to `chanIn` channel for incoming jobs
				for job := range chanIn {
					// Do the jobs
					filePath := filepath.Join(tempPath, job.FileName)
					content := randomString2(contentLength)
					err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

					// Construct the job's result, and send it to `chanOut`
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
