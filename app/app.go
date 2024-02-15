package app

import (
	"bufio"
	"fmt"
	"go_basic/db"
	"go_basic/ui"
	"log"
	"os"
	"strings"
)

func AddUrl() {
	for {
		args, _ := ui.ReadInput("Введите новую запись в формате <url описание тег1,тег2>")
		if len(args) != 3 {
			log.Println("формат аргументов не соответствует <url описание тег1,тег2>")
			continue
		}

		u := db.UrlData{
			Link:        args[0],
			Description: args[1],
			Tags:        args[2],
		}

		if err := db.SaveLink(u); err != nil {
			log.Printf("ошибка сохранения в базу данных (%s)\n", err)
			continue
		}

		log.Println("Данные о ссылке успешно сохранены")
		break
	}
}

func ListUrls() {
	urls := db.GetUrlList()
	ui.PrintUrls(urls)
}

func SearchUrls() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Выберите тип поиска: 'url' для поиска по URL, 'tags' для поиска по тегам")
	searchType, _ := reader.ReadString('\n')
	searchType = strings.TrimSpace(searchType) // Удаление символа новой строки

	if searchType != "url" && searchType != "tags" {
		fmt.Println("Неверный тип поиска. Пожалуйста, введите 'url' или 'tags'.")
		return
	}

	fmt.Println("Введите начало искомого слова:")
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query) // Удаление символа новой строки

	results, err := db.SearchUrls(searchType, query)
	if err != nil {
		fmt.Printf("Ошибка при выполнении поиска: %v\n", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("По вашему запросу ничего не найдено.")
		return
	}

	fmt.Println("Результаты поиска:")
	for _, url := range results {
		fmt.Printf("URL: %s, Описание: %s, Теги: %s\n", url.Link, url.Description, url.Tags)
	}
}

func DeleteUrl() {
	args, _ := ui.ReadInput("Введите URL для удаления:")
	if len(args) != 1 {
		log.Println("Необходимо ввести один URL")
		return
	}

	if err := db.DeleteUrl(args[0]); err != nil {
		log.Printf("Ошибка при удалении URL (%s)\n", err)
	} else {
		log.Println("URL успешно удален")
	}
}
