package main

import (
	"fmt"
	"go_basic/extractors"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePath := os.Args[1]

	fileName := extractors.FileName(filePath)
	nameWithoutExt, fileExt := extractors.FileExtension(fileName)

	fmt.Printf("filename: %s\n", nameWithoutExt)
	fmt.Printf("extension: %s\n", fileExt)
}
