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

	fullFileName := extractors.FileName(filePath)
	fileName, fileExt := extractors.FileExtension(fullFileName)

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
