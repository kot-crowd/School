from api_client import fetch_site_metrics
from parser import extract_metrics
from report_generator import generate_text_report
from export_report import export_metrics_to_csv

def main():
    site_url = input("Введите URL сайта конкурента: ").strip()
    if not site_url.startswith(("http://", "https://")):
        site_url = "https://" + site_url

    try:
        api_response = fetch_site_metrics(site_url)
        metrics = extract_metrics(api_response)
        report = generate_text_report(metrics)
        print("\n--- Отчёт по сайту ---\n")
        print(report)
        print("\n---------------------\n")

        # Экспорт в CSV для маркетолога
        csv_file = export_metrics_to_csv(metrics, "competitor_report.csv")
        print(f"Отчёт сохранён в файл: {csv_file}")
    except requests.exceptions.RequestException as e:
        print(f"Ошибка при запросе к API: {e}")
    except Exception as e:
        print(f"Произошла непредвиденная ошибка: {e}")

if __name__ == "__main__":
    main()
