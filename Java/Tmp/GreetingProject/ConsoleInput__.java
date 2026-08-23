package Java.Tmp.GreetingProject;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.UnsupportedEncodingException;

public class ConsoleInput__ implements AutoCloseable {
    private final BufferedReader reader;

    public ConsoleInput__() throws UnsupportedEncodingException {
        // Явно используем UTF-8 для чтения из консоли
        this.reader = new BufferedReader(new InputStreamReader(System.in, "UTF-8"));
    }

    public String readName() throws IOException {
        System.out.print("Введите ваше имя: ");
        return reader.readLine();
    }

    public int readAge() throws IOException {
        while (true) {
            try {
                System.out.print("Введите ваш возраст: ");
                String line = reader.readLine();
                int age = Integer.parseInt(line);
                if (age > 0) {
                    return age;
                } else {
                    System.out.println("Возраст должен быть положительным числом. Попробуйте снова.");
                }
            } catch (NumberFormatException e) {
                System.out.println("Ошибка: необходимо ввести целое число. Попробуйте снова.");
            }
        }
    }

    @Override
    public void close() throws IOException {
        reader.close();
    }
}
