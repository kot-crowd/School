package Java.GreetingProject;

public class GreetingFormatter {
    public static String format(String name, int age, String yearWord) {
        return String.format("Привет, %s! Тебе %d %s", name, age, yearWord);
    }
}