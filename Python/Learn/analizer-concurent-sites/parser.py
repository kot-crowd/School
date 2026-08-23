def extract_metrics(data: dict) -> dict:
    # Адаптируйте ключи под реальный ответ вашего API
    title = data.get("result", {}).get("title")
    meta_description = data.get("result", {}).get("meta_description")
    page_count = data.get("result", {}).get("page_count")

    return {
        "title": title,
        "meta_description": meta_description,
        "page_count": page_count
    }
