import sys

def concatenate_files(filenames):
    """
    Объединяет содержимое нескольких файлов в указанном порядке,
    обрабатывая возможные ошибки при открытии файлов.
    """
    content = ''
    if not filenames:
        print("Ошибка: не переданы имена файлов для объединения.")
        return

    for filename in filenames:
        try:
            with open(filename, 'r', encoding='utf-8') as f:
                content += f.read()

        except Exception as e:
            print(f"Произошла ошибка при работе с файлом '{filename}': {e}. Пропускаем.")

    print(content)

if __name__ == "__main__":
    # sys.argv[0] - это имя самого скрипта, поэтому начинаем с индекса 1
    file_list = sys.argv[1:]
    concatenate_files(file_list)