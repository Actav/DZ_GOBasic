package ui

import (
	"bufio"
	"fmt"
	"go_basic/db"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

func ReadInput(prompt string) ([]string, error) {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return strings.Fields(text), nil
}

func GetKeyPress() (rune, keyboard.Key, error) {
	return keyboard.GetKey()
}

func PrintUrls(urls []db.UrlData) {
	fmt.Printf("\nВсего добавленно urls: %d\n\n", len(urls))
	for i, url := range urls {
		fmt.Println(i+1, "\n---------------------------")
		fmt.Printf("Дата: %s\nURL: %s\nОписание: %s\nТеги: %s\n\n", url.CreateTime, url.Link, url.Description, url.Tags)
	}
}
