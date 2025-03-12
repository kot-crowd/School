package main

import (
	"encoding/json" // Импортируем пакет для работы с JSON
	"fmt"           // Импортируем пакет для форматирования вывода
	"os"            // Импортируем пакет для ввода-вывода

	//"io/ioutil"     // Импортируем пакет для ввода-вывода
	"log" // Импортируем пакет для логирования ошибок
)

// Структура для представления информации о группе студентов
type Group struct {
	ID       int       `json:"ID"`       // Идентификатор группы
	Number   string    `json:"Number"`   // Номер группы
	Year     int       `json:"Year"`     // Курс группы
	Students []Student `json:"Students"` // Массив студентов
}

// Структура для представления информации о студенте
type Student struct {
	LastName   string `json:"LastName"`   // Фамилия студента
	FirstName  string `json:"FirstName"`  // Имя студента
	MiddleName string `json:"MiddleName"` // Отчество студента
	Birthday   string `json:"Birthday"`   // Дата рождения студента
	Address    string `json:"Address"`    // Адрес студента
	Phone      string `json:"Phone"`      // Номер телефона студента
	Rating     []int  `json:"Rating"`     // Массив оценок
}

// Структура для представления результата
type Result struct {
	Average float64 `json:"Average"` // Среднее значение оценок
}

// Функция для расчитывания среднего количества оценок
func calculateAverageRating(students []Student) float64 {
	totalRatings := 0 // Переменная для хранения общей суммы оценок
	totalCount := 0   // Переменная для подсчета общего количества оценок

	// Итерируемся по каждому студенту
	for _, student := range students {
		totalRatings += sum(student.Rating) // Добавляем сумму оценок студента к общей
		totalCount += len(student.Rating)   // Увеличиваем общий счетчик оценок
	}

	// Проверяем, есть ли оценки, чтобы избежать деления на ноль
	if totalCount == 0 {
		return 0
	}

	// Возвращаем среднее значение
	return float64(totalRatings) / float64(totalCount)
}

// Функция для вычисления суммы оценок
func sum(ratings []int) int {
	total := 0
	for _, rating := range ratings {
		total += rating // Суммируем оценки
	}
	return total // Возвращаем общую сумму
}

// Главная функция
func main() {
	// Читаем файл students.json
	//data, err := ioutil.ReadFile("students.json")
	data, err := os.ReadFile("students.json")
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err) // Логируем ошибку, если не удалось прочитать файл
	}

	// Декодируем JSON данные в структуру Group
	var group Group
	if err := json.Unmarshal(data, &group); err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %v", err) // Логируем ошибку при декодировании
	}

	// Рассчитываем среднее количество оценок
	average := calculateAverageRating(group.Students)

	// Готовим результат для вывода
	result := Result{Average: average}

	// Кодируем результат в формат JSON с отступами
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Ошибка при кодировании JSON: %v", err) // Логируем ошибку при кодировании
	}

	// Выводим результат на стандартный вывод
	fmt.Println(string(output)) // Печатаем результат в формате JSON
}
