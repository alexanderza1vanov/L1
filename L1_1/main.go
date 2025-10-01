package main

import "fmt"

type Human struct {
	Name       string
	Age        int
	Address    string
	University string
}

// String() специальный метод в Go, который позволяет кастомизировать строковое представление структур
func (h Human) String() string {
	return fmt.Sprintf("Имя: %s\nВозраст: %d\nАдрес: %s\nУниверситет: %s ", h.Name, h.Age, h.Address, h.University)
}

type Action struct {
	Human
}

func main() {
	action := Action{Human{
		Name:       "Aleksandr",
		Age:        20,
		Address:    "Moscow",
		University: "RTU MIREA",
	}}

	fmt.Println(action) // Происходит неявный вызов String()

	//	Вывод:
	//	Имя: Aleksandr
	//	Возраст: 20
	//	Адрес: Moscow
	//	Университет: RTU MIREA

}
