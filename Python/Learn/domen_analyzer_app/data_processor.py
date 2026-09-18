"""Модуль преобразования сырых данных о домене в структурированный вид."""

from __future__ import annotations

from typing import Any, Optional


# Соответствие внутренних статусов RDAP и их отображаемых названий.
STATUS_LABELS: dict[str, str] = {
    "active": "REGISTERED",
    "registered": "REGISTERED",
    "delegated": "DELEGATED",
    "verified": "VERIFIED",
}

# Поля vCard, которые нас интересуют, и их внутренние ключи.
VCARD_FIELDS: dict[str, str] = {
    "fn": "organization",
    "country": "country",
    "taxpayer-id": "taxpayer_id",
}


def _parse_vcard(vcard_array: list[Any]) -> dict[str, str]:
    """Извлекает данные из vCard-массива RDAP."""
    result: dict[str, str] = {}
    if not vcard_array or len(vcard_array) < 2:
        return result

    for entry in vcard_array[1]:
        if not isinstance(entry, (list, tuple)) or len(entry) < 4:
            continue
        field_name = entry[0]
        value = entry[3]
        if field_name in VCARD_FIELDS:
            result[VCARD_FIELDS[field_name]] = value
    return result


def _find_entity(raw: dict[str, Any], role: str) -> dict[str, Any]:
    for entity in raw.get("entities", []) or []:
        if role in (entity.get("roles") or []):
            return entity
    return {}


def _get_event_date(events: list[dict[str, Any]], action: str) -> Optional[str]:
    for event in events or []:
        if event.get("eventAction") == action:
            return event.get("eventDate")
    return None


def _format_nameservers(nameservers: list[dict[str, Any]]) -> Optional[str]:
    if not nameservers:
        return None

    parts: list[str] = []
    for ns in nameservers:
        name = ns.get("ldhName") or ns.get("unicodeName")
        if not name:
            continue

        ip_data = ns.get("ipAddresses") or {}
        ips = list(ip_data.get("v4", []) or []) + list(ip_data.get("v6", []) or [])

        parts.append(f"{name} {', '.join(ips)}" if ips else name)

    return ", ".join(parts) if parts else None


def _format_status(statuses: list[str]) -> Optional[str]:
    if not statuses:
        return None
    return ", ".join(STATUS_LABELS.get(s.lower(), s.upper()) for s in statuses)


class DomainDataProcessor:
    """Преобразует сырые данные RDAP в структурированный словарь."""

    def process(self, raw: dict[str, Any]) -> dict[str, Any]:
        registrant = _find_entity(raw, "registrant")
        registrar = _find_entity(raw, "registrar")

        registrant_vcard = _parse_vcard(registrant.get("vcardArray", []))
        registrar_vcard = _parse_vcard(registrar.get("vcardArray", []))

        events = raw.get("events", []) or []
        domain_name = (raw.get("ldhName") or raw.get("unicodeName") or "").upper() or None

        return {
            "доменное имя": domain_name,
            "владелец или организация": registrant_vcard.get("organization"),
            "регистратор": registrar_vcard.get("organization"),
            "дата регистрации": _get_event_date(events, "registration"),
            "дата истечения": _get_event_date(events, "expiration"),
            "дата последнего обновления": _get_event_date(events, "last changed"),
            "серверы имён": _format_nameservers(raw.get("nameservers", []) or []),
            "статус домена": _format_status(raw.get("status", []) or []),
            "страна регистранта": registrant_vcard.get("country"),
            # taxpayer-id → ИНН регистранта (идентификационный номер налогоплательщика).
            "ИНН регистранта": registrant_vcard.get("taxpayer_id"),
        }


# --- Демонстрация ----------------------------------------------------------

if __name__ == "__main__":
    from api_client import MockDomainApiClient

    client = MockDomainApiClient()
    processor = DomainDataProcessor()

    data = client.get_domain_info("yandex.ru")
    result = processor.process(data)

    for key, value in result.items():
        print(f"{key}: {value}")