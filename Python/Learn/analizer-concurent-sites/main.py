from api_client import fetch_site_metrics
from parser import extract_metrics
from report_generator import generate_text_report
import requests

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
    except requests.exceptions.RequestException as e:
        print(f"Ошибка при запросе к API: {e}")
    except Exception as e:
        print(f"Произошла непредвиденная ошибка: {e}")

if __name__ == "__main__":
    main()
