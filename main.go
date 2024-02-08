package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите предложение для анализа:")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence) // Удаляем символ новой строки

	letterCounts := make(map[rune]int)
	var totalLetters int

	// Подсчет встречаемости каждой буквы и общего количества букв
	for _, letter := range sentence {
		if letter != ' ' && letter != '\n' { // Пропускаем пробелы и символы новой строки
			letterCounts[letter]++
			totalLetters++
		}
	}

	// Вывод количества встреч каждой буквы и частоты в процентах
	for letter, count := range letterCounts {
		frequency := float64(count) / float64(totalLetters)
		fmt.Printf("%c - %d %.2f\n", letter, count, frequency)
	}
}
