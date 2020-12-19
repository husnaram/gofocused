package pipeline

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var localTempPath = filepath.Join(os.Getenv("TEMP"), "ch-A59-pipeline-temp")

func MainProcess() {
	log.Println("Start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(), "seconds")
}

func proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(localTempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		counterTotal++

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		sum := fmt.Sprintf("%x", md5.Sum(buf))

		destinationPath := filepath.Join(localTempPath, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}
