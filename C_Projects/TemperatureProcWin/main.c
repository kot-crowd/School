/*Для мониторинга температуры процессора на Windows можно использовать библиотеку Windows API. Ниже представлен пример кода на языке C, который использует
библиотеку Windows.h для извлечения информации о температуре процессора. Учтите, что для получения данных о температуре может потребоваться сторонняя библиотека
или инструмент, так как стандартный Windows API не предоставляет доступа к температурным датчикам напрямую.

Пример кода на C с использованием WMI
Для доступа к информации о температуре процессора через WMI можно использовать библиотеку wbemidl.h. В Windows требуется подключить соответствующие библиотеки и заголовки.

Сохраните приведенный ниже код в файл, например, cpu_temp_monitor.c.*/
#include <stdio.h>
#include <windows.h>
#include <wbemidl.h>
#include <comdef.h>

#pragma comment(lib, "wbemuuid.lib")

void print_error(const char *msg)
{
    fprintf(stderr, "%s: %s\n", msg, _com_error(GetLastError()).ErrorMessage());
}

void QueryCPUTemperature()
{
    HRESULT hres;

    // Инициализация COM.
    hres = CoInitializeEx(0, COINIT_MULTITHREADED);
    if (FAILED(hres))
    {
        print_error("Failed to initialize COM");
        return;
    }

    hres = CoInitializeSecurity(
        NULL,
        -1,
        NULL,
        NULL,
        RPC_C_AUTHN_LEVEL_DEFAULT,
        RPC_C_IMP_LEVEL_IMPERSONATE,
        NULL,
        EOAC_NONE,
        NULL);

    if (FAILED(hres))
    {
        print_error("Failed to initialize security");
        CoUninitialize();
        return;
    }

    // Создание объекта WMI.
    IWbemLocator *pLoc = NULL;
    hres = CoCreateInstance(
        CLSID_WbemLocator,
        0,
        CLSCTX_INPROC_SERVER,
        IID_IWbemLocator, (LPVOID *)&pLoc);

    if (FAILED(hres))
    {
        print_error("Failed to create IWbemLocator object");
        CoUninitialize();
        return;
    }

    // Подключение к WMI.
    IWbemServices *pSvc = NULL;
    hres = pLoc->ConnectServer(
        _bstr_t(L"ROOT\\CIMV2"),
        NULL, NULL, 0,
        NULL, 0, 0,
        &pSvc);

    if (FAILED(hres))
    {
        print_error("Could not connect to WMI");
        pLoc->Release();
        CoUninitialize();
        return;
    }

    // Установка безопасности для WMI.
    hres = CoInitializeSecurity(
        NULL,
        -1,
        NULL,
        NULL,
        RPC_C_AUTHN_LEVEL_DEFAULT,
        RPC_C_IMP_LEVEL_IMPERSONATE,
        NULL,
        EOAC_NONE,
        NULL);

    if (FAILED(hres))
    {
        print_error("Could not initialize security");
        pSvc->Release();
        pLoc->Release();
        CoUninitialize();
        return;
    }

    // Запрос данных о температуре.
    IEnumWbemClassObject *pEnumerator = NULL;
    hres = pSvc->ExecQuery(
        bstr_t("WQL"),
        bstr_t("SELECT CurrentTemperature FROM Win32_Temperature"),
        WBEM_FLAG_FORWARD_ONLY | WBEM_FLAG_RETURN_IMMEDIATELY,
        NULL,
        &pEnumerator);

    if (SUCCEEDED(hres))
    {
        IWbemClassObject *pclsObj = NULL;
        ULONG uReturn = 0;

        while (pEnumerator)
        {
            HRESULT hr = pEnumerator->Next(WBEM_INFINITE, 1, &pclsObj, &uReturn);
            if (0 == uReturn)
                break;

            // Получение значения температуры.
            VARIANT vtProp;
            hr = pclsObj->Get(L"CurrentTemperature", 0, &vtProp, 0, 0);
            printf("Температура процессора: %.2f °C\n", (vtProp.vt != VT_NULL) ? (vtProp.floatVal / 10.0f - 273.15f) : -1.0f);
            VariantClear(&vtProp);
            pclsObj->Release();
        }
    }
    else
    {
        print_error("Query for temperature failed");
    }

    // Освобождение ресурсов
    pSvc->Release();
    pLoc->Release();
    pEnumerator->Release();
    CoUninitialize();
}

int main()
{
    QueryCPUTemperature();
    return 0;
}
/*Как компилировать и запускать код:
Сохраните код в файл, например, cpu_temp_monitor.c.
Используйте компилятор, например, MinGW или Microsoft Visual Studio. Если используете MinGW, можно использовать следующую команду:
gcc -o cpu_temp_monitor cpu_temp_monitor.c -lole32 -loleaut32 -lwbemuuid
Запустите программу.
./cpu_temp_monitor
Примечания:
CPUTemperature может не отображаться, если ваш компьютер не поддерживает WMI для получения информации о температуре.
В некоторых системах может потребоваться установка дополнительного ПО или драйверов для доступа к температурным датчикам.
Код прост и требует дополнительной обработки ошибок и улучшений для использования в производственной среде.*/