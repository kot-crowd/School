package main

import (
	"encoding/json" // Пакет для работы с JSON
	"fmt"           // Пакет для форматированного ввода и вывода
	"log"           // Пакет для логирования ошибок
	"os"            // Пакет для работы с файлами
)

// Определяем структуру для элементов из файла структуры
type StructureItem struct {
	ID string `json:"id"` // Поле id для сопоставления
	// При необходимости добавьте другие поля
}

// Определяем структуру для элементов из файла данных
type DataItem struct {
	GlobalID string `json:"global_id"` // Поле global_id, значение которого необходимо суммировать
	ID       string `json:"id"`        // Поле id для сопоставления с элементами структуры
	// При необходимости добавьте другие поля
}

func main() {
	// Определяем пути к файлам
	structureFilePath := "/home/kt/GoProjects/StructureInFiles/structure-20190514T0000.json"
	dataFilePath := "/home/kt/GoProjects/StructureInFiles/data-20190514T0100.json"

	// Читаем содержимое файла структуры
	structureData, err := os.ReadFile(structureFilePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла структуры: %v", err) // Логируем ошибку, если чтение файла не удалось
	}

	// Читаем содержимое файла данных
	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла данных: %v", err) // Логируем ошибку, если чтение файла не удалось
	}

	// Объявляем срез для хранения элементов структуры
	var structureItems []StructureItem
	// Разбираем JSON-данные из файла структуры в срез
	if err := json.Unmarshal(structureData, &structureItems); err != nil {
		log.Fatalf("Ошибка при разборе файла структуры: %v", err) // Логируем ошибку, если разбор не удался
	}

	// Объявляем срез для хранения элементов данных
	var dataItems []DataItem
	// Разбираем JSON-данные из файла данных в срез
	if err := json.Unmarshal(data, &dataItems); err != nil {
		log.Fatalf("Ошибка при разборе файла данных: %v", err) // Логируем ошибку, если разбор не удался
	}

	// Создаем множество ID из структуры для сопоставления
	structureIDMap := make(map[string]struct{})
	for _, item := range structureItems {
		structureIDMap[item.ID] = struct{}{}
	}

	// Вычисляем сумму полей global_id
	var sum int
	for _, item := range dataItems {
		if _, exists := structureIDMap[item.ID]; exists {
			var id int
			if _, err := fmt.Sscanf(item.GlobalID, "%d", &id); err == nil {
				sum += id
			} else {
				log.Printf("Ошибка при парсинге global_id: %v", err)
			}
		}
	}

	// Выводим результат
	fmt.Printf("Сумма полей global_id всех соответствующих элементов: %d\n", sum)
}
