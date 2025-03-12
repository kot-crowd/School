#Шаги для создания скрипта:
#Убедитесь, что библиотеки установлены. Выполните следующие команды в терминале:
#sudo apt install tesseract-ocr
#pip install pytesseract Pillow
#Создайте Python-скрипт. Например, назовем его ocr_images.py и вставьте следующий код:
#Описание кода:
#Импорт библиотек: Импортируются необходимые модули для работы с изображениями и распознаванием текста.
#Настройка Tesseract: Укажите путь к исполняемому файлу Tesseract. В большинстве случаев он находится по адресу /usr/bin/tesseract, но вы можете проверить это с 
#помощью команды which tesseract в терминале.
#Указание папки с изображениями: Задайте путь к папке, в которой находятся ваши изображения (например, 'images/').
#Обработка изображений: Скрипт проверяет все файлы в указанной папке. Если файл является изображением, он открывается, и на его основе выполняется OCR с помощью
#pytesseract.
#Сохранение результата: Распознанный текст сохраняется в текстовом файле output.txt.
#Запуск скрипта
#Сохраните скрипт и запустите его в терминале:

#python ocr_images.py
#Убедитесь, что все изображения находятся в указанной папке, и выполните скрипт для распознавания текста. Вы получите результат в текстовом файле output.txt.

import pytesseract
from PIL import Image
import os

# Укажите, где установлен Tesseract (в зависимости от вашей установки)
pytesseract.pytesseract.tesseract_cmd = '/usr/bin/tesseract'  # Убедитесь, что это правильно

# Путь к папке с изображениями
image_folder = 'путь_к_вашей_папке_с_изображениями'  # Укажите путь к папке
output_file = 'output.txt'  # Имя файла для сохранения распознанного текста

# Список для хранения распознанного текста
recognized_texts = []

# Обход всех файлов в папке
for filename in os.listdir(image_folder):
    # Проверяем, является ли файл изображением (можно добавить больше типов)
    if filename.endswith(('.png', '.jpg', '.jpeg', '.gif', '.bmp')):
        image_path = os.path.join(image_folder, filename)
        try:
            # Открываем изображение и выполняем OCR
            image = Image.open(image_path)
            text = pytesseract.image_to_string(image, lang='rus')  # Укажите нужный язык
            recognized_texts.append(f"--- {filename} ---\n{text}\n")  # Добавляем текст к списку
        except Exception as e:
            print(f"Ошибка при обработке файла {filename}: {e}")

# Сохраняем распознанный текст в файл
with open(output_file, 'w', encoding='utf-8') as f:
    f.writelines(recognized_texts)

print(f"Распознанный текст сохранен в {output_file}")