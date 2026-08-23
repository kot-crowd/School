package Java.GreetingProject;

import java.util.InputMismatchException;
import java.util.Scanner;

public class ConsoleInput {
    public static String readName(Scanner scanner) {
        System.out.print("Введите ваше имя: ");
        return scanner.nextLine();
    }

    public static int readAge(Scanner scanner) {
        while (true) {
            try {
                System.out.print("Введите ваш возраст: ");
                int age = scanner.nextInt();
                if (age > 0) {
                    return age;
                } else {
                    System.out.println("Возраст должен быть положительным числом. Попробуйте снова.");
                }
            } catch (InputMismatchException e) {
                System.out.println("Ошибка: необходимо ввести целое число. Попробуйте снова.");
                scanner.nextLine(); // очистка буфера
            }
        }
    }
}