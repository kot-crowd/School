# export_report.py
import csv
from typing import Dict, Any

def export_metrics_to_csv(metrics: Dict[str, Any], filename: str = "competitor_report.csv") -> str:
    """
    Экспортирует метрики в CSV-файл для маркетолога.

    Параметры:
        metrics (dict): словарь с метриками (title, meta_description, page_count).
        filename (str): имя выходного CSV-файла.

    Возвращает:
        str: путь к сохранённому файлу.
    """
    # Определяем порядок колонок и их читаемые названия
    fieldnames = ["title", "meta_description", "page_count"]
    header_row = {
        "title": "Заголовок страницы",
        "meta_description": "Мета‑описание",
        "page_count": "Количество страниц"
    }

    # Формируем одну строку данных из метрик
    row_data = {
        key: metrics.get(key) if metrics.get(key) is not None else "Не найдено"
        for key in fieldnames
    }

    with open(filename, mode="w", newline="", encoding="utf-8") as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)

        # Записываем заголовок (читаемые названия колонок)
        writer.writerow(header_row)
        # Записываем строку с данными
        writer.writerow(row_data)

    return filename
