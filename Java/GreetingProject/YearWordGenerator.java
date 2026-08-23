package Java.GreetingProject;

public class YearWordGenerator {
    public static String getYearWord(int years) {
        int lastDigit = years % 10;
        int lastTwo = years % 100;

        if (11 <= lastTwo && lastTwo <= 14) {
            return "лет";
        } else if (lastDigit == 1) {
            return "год";
        } else if (2 <= lastDigit && lastDigit <= 4) {
            return "года";
        } else {
            return "лет";
        }
    }
}
