package main

import "fmt"

type stringMap = map[string]string

func main() {
	fmt.Println("__Заметки__")
	bookmarks := stringMap{}
Menu:
	for {
		numMenu := menu()

		switch numMenu {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		default:
			break Menu
		}
	}
}

func deleteBookmark(bookmarks stringMap) {
	var name string
	fmt.Print("Ввыдите имя закладки: ")
	fmt.Scan(&name)
	delete(bookmarks, name)
}

func addBookmark(bookmarks stringMap) {
	var name, text string
	fmt.Print("Введите имя закладки: ")
	fmt.Scan(&name)
	fmt.Print("Введите текст закладки: ")
	fmt.Scan(&text)
	bookmarks[name] = text
}

func printBookmarks(bookmarks stringMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Закладки пустые")
	}
	for name, desc := range bookmarks {
		fmt.Printf("%v: %v\n", name, desc)
	}
}

func menu() int {
	var num int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладку")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Print("> ")
	fmt.Scan(&num)
	return num
}
