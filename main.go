package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {

	files := os.Args[1:]
	wg := sync.WaitGroup{}

	for _, file := range files {
		output := file + ".zip"
		wg.Add(1)
		go ZipFile(&wg, output, file)
	}
	wg.Wait()
	fmt.Print("done")
}

func ZipFile(wg *sync.WaitGroup, filename string, file string)  {
	defer wg.Done()
	newZipFile, err := os.Create(filename)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()



		zipfile, err := os.Open(file)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		defer zipfile.Close()

		info, err := zipfile.Stat()
		if err != nil {
			log.Printf("%v", err)
			return
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Printf("%v", err)
			return
		}

		header.Name = file

		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		if _, err = io.Copy(writer, zipfile); err != nil {
			log.Printf("%v", err)
			return
		}
	log.Println("Zipped File: " + filename)
}

