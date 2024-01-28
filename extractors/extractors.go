package extractors

import "strings"

// FileName извлекает имя файла из полного пути.
func FileName(filePath string) string {
	if i := strings.LastIndexByte(filePath, '/'); i != -1 {
		return filePath[i+1:]
	}

	return filePath
}

// FileExtension извлекает расширение файла из имени файла.
func FileExtension(fileName string) (string, string) {
	if i := strings.LastIndexByte(fileName, '.'); i != -1 {
		return fileName[:i], fileName[i+1:]
	}

	return fileName, ""
}
