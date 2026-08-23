def check_grade(score):
    if score >= 90:
        return "Отлично"
    elif score >= 80:
        return "Хорошо"
    elif score >= 70:
        return "Удовлетворительно"
    else:
        return "Неудовлетворительно"

def calculate_sum(numbers):
    total = 0
    for num in numbers:
        total = total + num
    return total

print("Программа для проверки оценок")
print("Введите баллы студента:")

math_score = int(input("Математика: "))
physics_score = int(input("Физика: "))
chemistry_score = int(input("Химия: "))

print("Результаты:")
print("Математика:", check_grade(math_score))
print("Физика:", check_grade(physics_score))
print("Химия:", check_grade(chemistry_score))

scores = [math_score, physics_score, chemistry_score]
total = calculate_sum(scores)
average = total / 3

print("Общий балл:", total)
print("Средний балл:", average)

if average >= 80:
    print("Студент переведён на следующий курс")
else:
    print("Студенту нужны дополнительные занятия")