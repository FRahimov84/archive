package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"os"
	"sync"
)

var (
	competitive  = flag.Bool("competitive", false, "competitive archivator")
	consistently = flag.Bool("consistently", false, "consistently archivator")
)

func main() {
	flag.Parse()
	files := os.Args[2:]
	if *consistently {
		consistentlyArchiver(files)
		return
	}

	if *competitive {
		competitiveArchiver(files)
		return
	}

	log.Print("incorrect flag")
}

func competitiveArchiver(files []string) {
	wg := sync.WaitGroup{}

	for _, file := range files {
		output := "competitive/" + file + ".zip"
		wg.Add(1)
		go func() {
			defer wg.Done()
			ZipFile(output, file)
		}()
	}
	wg.Wait()
}

func consistentlyArchiver(files []string) {

	for _, file := range files {
		output := "consistently/" + file + ".zip"
		ZipFile(output, file)
	}
}

func ZipFile(filename string, file string) {

	newZipFile, err := os.Create(filename)
	if err != nil {
		log.Printf("can't create file %s err: %v", filename, err)
		return
	}
	defer func() {
		err = newZipFile.Close()
		if err != nil {
			log.Printf("can't close newZipFile err: %v", err)
		}
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		err = zipWriter.Close()
		if err != nil {
			log.Printf("can't close zipWriter err: %v", err)
		}
	}()
	zipfile, err := os.Open(file)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer func() {
		err = zipfile.Close()
		if err != nil {
			log.Printf("can't close zipfile err: %v", err)
		}
	}()
	info, err := zipfile.Stat()
	if err != nil {
		log.Printf("can't get zipfile stat err: %v", err)
		return
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		log.Printf("can't get FileInfoHeader err: %v", err)
		return
	}
	header.Name = file

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		log.Printf("can't create zipWriter err: %v", err)
		return
	}
	if _, err = io.Copy(writer, zipfile); err != nil {
		log.Printf("can't copy ti zipfile err: %v", err)
		return
	}
	//log.Println("Zipped File: " + filename)
}

