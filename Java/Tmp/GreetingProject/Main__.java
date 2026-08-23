package Java.Tmp.GreetingProject;

import java.io.PrintStream;
import java.io.UnsupportedEncodingException;

import Java.GreetingProject.AgeValidator;
import Java.GreetingProject.GreetingFormatter;
import Java.GreetingProject.YearWordGenerator;

public class Main__ {
    public static void main(String[] args) {
        try {
            // Устанавливаем кодировку для вывода UTF-8
            System.setOut(new PrintStream(System.out, true, "UTF-8"));
        } catch (UnsupportedEncodingException e) {
            System.err.println("Ошибка установки кодировки вывода: " + e.getMessage());
        }

        try (ConsoleInput__ consoleInput = new ConsoleInput__()) {
            String name = consoleInput.readName();
            int age = consoleInput.readAge();

            if (!AgeValidator.isAdult(age)) {
                System.out.println("Доступ запрещён. Вам должно быть 18 лет или больше.");
                return;
            }

            String yearWord = YearWordGenerator.getYearWord(age);
            String greeting = GreetingFormatter.format(name, age, yearWord);
            System.out.println(greeting);
        } catch (Exception e) {
            System.err.println("Ошибка: " + e.getMessage());
        }
    }
}