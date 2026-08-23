import requests
from config import API_URL, TIMEOUT_SECONDS, HEADERS

def fetch_site_metrics(site_url: str):
    payload = {"url": site_url}
    response = requests.post(API_URL, json=payload, headers=HEADERS, timeout=TIMEOUT_SECONDS)
    response.raise_for_status()  # выбросит HTTPError при неудачном статусе
    return response.json()
