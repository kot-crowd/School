def remove_comments(filename_source):
    try:
        with open(filename_source, 'r', encoding='utf-8') as f:
            lines = f.readlines()
            for i, line in enumerate(lines):
                if "#" in lines[i]:
                    lines[i] = lines[i].split('#')[0]  + "#\n"
        return lines
    except Exception as e:
                print(f"Ошибка при работе с файлом '{filename_source}': {e}.")


def new_file(filename_solution, lines):
    try:
        with open(filename_solution, 'w', encoding='utf-8') as f:
            f.writelines(lines)
    except Exception as e:
                print(f"Ошибка при работе с файлом '{filename_solution}': {e}.")

filename_source = ''
lines_edit = []
while not filename_source.endswith(".py"):
    filename_source = input("Введите имя файла, с кодом на языке Python,\
который будем редактировать. Обратите внимание, у файла должно быть расширение 'py'\n")
    
lines_edit = remove_comments(filename_source)

filename_solution = ''
while not filename_solution.endswith(".py"):
    filename_solution = input("Введите имя файла, c кодом на языке Python,\
в который внесём результат. Обратите внимание, у файла должно быть расширение '.py'\n")

new_file(filename_solution, lines_edit)