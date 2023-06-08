package main

import (
	"encoding/json"
	"fmt"
)

// Объявляем тип user
type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func main() {
	// Объявляем и заполняем структуру newUser типа user
	newUser := user{
		Name:     "Alex",
		Email:    "email@email.email",
		Status:   true,
		Language: []byte("ru"),
	}
	// Объявляем структуру data и передаём в неё значение структуры newUser
	data, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}
	// Меняем значение параметра language структуры newUser
	newUser.Language = []byte("en")
	// Передаём значение структуры data в структуру newUser
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		panic(err)
	}
	// Выводим значение параметра Language структуры newUser
	fmt.Println(string(newUser.Language))
}
