def generate_text_report(metrics: dict) -> str:
    lines = []
    if metrics["title"] is not None:
        lines.append(f"Заголовок страницы: {metrics['title']}")
    else:
        lines.append("Заголовок страницы: не найден")

    if metrics["meta_description"] is not None:
        lines.append(f"Мета‑описание: {metrics['meta_description']}")
    else:
        lines.append("Мета‑описание: не найдено")

    if metrics["page_count"] is not None:
        lines.append(f"Количество страниц: {metrics['page_count']}")
    else:
        lines.append("Количество страниц: не найдено")

    return "\n".join(lines)
