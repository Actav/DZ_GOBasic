package main

import (
	"fmt"
	"go_basic/app"
	"go_basic/db"
	"go_basic/ui"

	"github.com/eiannone/keyboard"
)

func main() {
	db.Init()
	defer db.DB.Close()

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")

	// Карта команд
	commands := map[rune]func(){
		'n': app.AddUrl,
		'l': app.ListUrls,
		'd': app.DeleteUrl,
		's': app.SearchUrls,
	}

	for {
		char, key, err := ui.GetKeyPress()
		if err != nil {
			panic(err)
		}

		if action, exists := commands[char]; exists {
			action()
		} else if key == keyboard.KeyEsc {
			break
		}
	}
}
