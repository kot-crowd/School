"""Модуль для получения сырых данных о доменных именах.

Содержит:
    * DomainApiClient     — базовый интерфейс клиента;
    * HttpDomainApiClient — реальный HTTP-клиент (RDAP-совместимый API);
    * MockDomainApiClient — мок-клиент для тестов и демонстрации.
"""

from __future__ import annotations

import json
import urllib.error
import urllib.request
from typing import Any


class DomainApiError(Exception):
    """Исключение при обращении к API регистратора."""


class DomainApiClient:
    """Базовый интерфейс клиента API."""

    def get_domain_info(self, domain: str) -> dict[str, Any]:
        raise NotImplementedError


class HttpDomainApiClient(DomainApiClient):
    """Клиент, обращающийся к RDAP-совместимому HTTP API."""

    def __init__(self, base_url: str = "https://rdap.example.org", timeout: int = 10) -> None:
        self.base_url = base_url.rstrip("/")
        self.timeout = timeout

    def get_domain_info(self, domain: str) -> dict[str, Any]:
        url = f"{self.base_url}/domain/{domain.strip().lower()}"
        request = urllib.request.Request(
            url,
            headers={
                "Accept": "application/rdap+json, application/json",
                "User-Agent": "DomainInfoClient/1.0",
            },
        )
        try:
            with urllib.request.urlopen(request, timeout=self.timeout) as response:
                payload = response.read().decode("utf-8")
        except urllib.error.URLError as exc:
            raise DomainApiError(
                f"Не удалось получить данные для домена {domain}: {exc}"
            ) from exc

        try:
            return json.loads(payload)
        except json.JSONDecodeError as exc:
            raise DomainApiError(f"Некорректный JSON в ответе API: {exc}") from exc


# --- Демонстрационная «база» для мок-клиента -------------------------------

_MOCK_DATABASE: dict[str, dict[str, Any]] = {
    "yandex.ru": {
        "ldhName": "yandex.ru",
        "status": ["active", "delegated", "verified"],
        "events": [
            {"eventAction": "registration", "eventDate": "1997-09-23T09:45:07Z"},
            {"eventAction": "expiration", "eventDate": "2026-09-30T21:00:00Z"},
        ],
        "entities": [
            {
                "roles": ["registrant"],
                "vcardArray": [
                    "vcard",
                    [
                        ["version", {}, "text", "4.0"],
                        ["fn", {}, "text", "Yandex LLC"],
                        # ИНН организации в формате vCard (RDAP-расширение).
                        ["taxpayer-id", {}, "text", "7736207543"],
                    ],
                ],
            },
            {
                "roles": ["registrar"],
                "vcardArray": [
                    "vcard",
                    [
                        ["version", {}, "text", "4.0"],
                        ["fn", {}, "text", "RU-CENTER-RU"],
                    ],
                ],
            },
        ],
        "nameservers": [
            {
                "ldhName": "ns1.yandex.ru.",
                "ipAddresses": {"v4": ["213.180.193.1"], "v6": ["2a02:6b8::1"]},
            },
            {
                "ldhName": "ns2.yandex.ru.",
                "ipAddresses": {"v4": ["93.158.134.1"], "v6": ["2a02:6b8::1"]},
            },
        ],
    },
}


class MockDomainApiClient(DomainApiClient):
    """Мок-клиент: возвращает заранее подготовленные данные."""

    def get_domain_info(self, domain: str) -> dict[str, Any]:
        key = domain.strip().lower()
        if key not in _MOCK_DATABASE:
            raise DomainApiError(f"Нет данных для домена {domain}")
        # Глубокая копия, чтобы вызывающий код не мог испортить «базу».
        return json.loads(json.dumps(_MOCK_DATABASE[key]))