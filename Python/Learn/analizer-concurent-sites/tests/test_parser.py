from parser import extract_metrics

def test_extract_metrics():
    data = {"result": {"title": "Тест", "meta_description": "Описание", "page_count": 10}}
    metrics = extract_metrics(data)
    assert metrics["title"] == "Тест"
    assert metrics["meta_description"] == "Описание"
    assert metrics["page_count"] == 10

