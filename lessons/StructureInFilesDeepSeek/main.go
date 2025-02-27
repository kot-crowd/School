package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type StructureItem struct {
	ID string `json:"id"`
}

type DataItem struct {
	GlobalID int    `json:"global_id"`
	ID       string `json:"id"`
}

type StructureWrapper struct {
	Items []StructureItem `json:"items"`
}

func main() {
	structureFilePath := "/home/kt/GoProjects/StructureInFilesDeepSeek/structure-20190514T0000.json"
	dataFilePath := "/home/kt/GoProjects/StructureInFilesDeepSeek/data-20190514T0100.json"

	structureData, err := os.ReadFile(structureFilePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла структуры: %v", err)
	}

	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла данных: %v", err)
	}

	var wrapper StructureWrapper
	if err := json.Unmarshal(structureData, &wrapper); err != nil {
		log.Fatalf("Ошибка при разборе файла структуры: %v", err)
	}

	structureItems := wrapper.Items

	var dataItems []DataItem
	if err := json.Unmarshal(data, &dataItems); err != nil {
		log.Fatalf("Ошибка при разборе файла данных: %v", err)
	}

	structureMap := make(map[string]bool)
	for _, item := range structureItems {
		structureMap[item.ID] = true
	}

	totalGlobalID := 0

	for _, item := range dataItems {
		if _, exists := structureMap[item.ID]; exists {
			totalGlobalID += item.GlobalID
		}
	}

	fmt.Printf("Сумма полей global_id всех соответствующих элементов: %d\n", totalGlobalID)
}
