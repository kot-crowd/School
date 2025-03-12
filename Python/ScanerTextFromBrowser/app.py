#Шаги для создания приложения:
#Установите необходимые библиотеки. Выполните следующие команды в терминале:
#pip install Flask pytesseract Pillow
#Установите Tesseract OCR. Скачайте и установите его с https://github.com/tesseract-ocr/tesseract. После установки убедитесь, что путь к Tesseract правильно указан в вашем коде.
#Создайте файл app.py и вставьте следующий код:
#Создайте файл templates/upload.html для интерфейса загрузки
#Объяснение кода:
#Создание веб-сервера на Flask. Мы создаем Flask-приложение и определяем два маршрута:
#GET для отображения формы загрузки.
#POST для обработки загруженного изображения.
#Загрузка файла. Когда пользователь отправляет изображение, файл сохраняется в указанную директорию и обрабатывается с использованием Pytesseract.
#Распознавание текста. Используется метод pytesseract.image_to_string() для извлечения текста, который затем возвращается пользователю в виде HTML.
#Запуск приложения:
#Сохраните оба файла.
#Выполните команду в терминале для запуска приложения:
#python app.py
#Откройте ваш веб-браузер и перейдите по адресу http://127.0.0.1:5000/.
#Теперь вы можете загрузить изображение, и приложение распознает текст с него и отобразит его на странице.

from flask import Flask, request, render_template
import pytesseract
from PIL import Image
import os

# Указать, где установлен Tesseract
pytesseract.pytesseract.tesseract_cmd = '/usr/bin/tesseract'  # Убедитесь, что это правильный путь

app = Flask(__name__)

# Путь к папке для хранения загруженных изображений
UPLOAD_FOLDER = 'uploads'
if not os.path.exists(UPLOAD_FOLDER):
    os.makedirs(UPLOAD_FOLDER)

@app.route('/')
def upload_form():
    return render_template('upload.html')

@app.route('/', methods=['POST'])
def upload_image():
    if 'image' not in request.files:
        return "Нет файла для загрузки!"

    file = request.files['image']

    if file.filename == '':
        return "Файл не выбран!"

    if file:
        # Сохраняем файл
        file_path = os.path.join(UPLOAD_FOLDER, file.filename)
        file.save(file_path)
        
        # Открываем изображение и выполняем OCR
        image = Image.open(file_path)
        text = pytesseract.image_to_string(image, lang='rus')  # Измените 'rus', если используете другой язык

        return f"Распознанный текст: <pre>{text}</pre>"

if __name__ == '__main__':
    app.run(debug=True)