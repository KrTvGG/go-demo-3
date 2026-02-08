package main

import "fmt"

func main() {
	m := map[string]string {
		"PurpleShcool": "purpleschool.ru",
	}

	fmt.Println(m)
	fmt.Println(m["PurpleShcool"])
	m["PurpleShcool"] = "https://purpleschool.ru"
	fmt.Println(m)
	m["Google"] = "https://google.com"
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	delete(m, "Y")
	fmt.Println(m["Y"])
	fmt.Println(m)
}
