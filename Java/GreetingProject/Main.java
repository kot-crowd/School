package Java.GreetingProject;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        System.setProperty("file.encoding", "UTF-8");
        try (Scanner scanner = new Scanner(System.in)) {
            // 1. Получить имя
            String name = ConsoleInput.readName(scanner);
            // 2. Получить возраст с проверкой положительности
            int age = ConsoleInput.readAge(scanner);
            // 3. Проверить возрастное ограничение
            if (!AgeValidator.isAdult(age)) {
                System.out.println("Доступ запрещён. Вам должно быть 18 лет или больше.");
                return;
            }
            // 4. Получить правильное слово "год/года/лет"
            String yearWord = YearWordGenerator.getYearWord(age);
            // 5. Сформировать и вывести приветствие
            String greeting = GreetingFormatter.format(name, age, yearWord);
            System.out.println(greeting);
        }
    }
}