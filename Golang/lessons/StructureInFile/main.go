package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Определяем структуру, соответствующую JSON данным
type Item struct {
	GlobalID int64 `json:"global_id"`
}

func main() {
	// Чтение данных из файла
	data, err := os.ReadFile("data-20190514T0100.json")
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	// Декодирование JSON данных
	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("Ошибка декодирования JSON: %v", err)
	}

	// Суммирование глобальных идентификаторов
	var sum int64
	for _, item := range items {
		sum += item.GlobalID
	}

	// Вывод результата
	fmt.Printf("Сумма полей global_id: %d\n", sum)
}
