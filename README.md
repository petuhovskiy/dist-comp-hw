# dist-comp-hw

## Что это?

Это ДЗ по курсу Распределенных вычислений. Этот репозиторий содержит реализацию скелета интернет магазина,
в котором доступна логика работы с товарами посредством REST API, реализованы все 5 методов:

1) Создавать товар
2) Удалять товар
3) Выводить список товаров (за пагинацию отдельный плюс)
4) Выводить отдельный товар
5) Редактировать товар

Также реализовано:

- [X] Хранение товаров в базе данных PostgreSQL
- [X] Версионирование API (пока что только версия 1)
- [X] Логирование запросов
- [X] Swagger UI доступен по адресу /swagger

Также есть сервис авторизации.

## Графическое описание архитектуры

Смотреть [здесь](https://app.diagrams.net?lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1#R7VnbktsoEP0aP9olCV3sR18zW5WtuNZVm80jlrBNIgsH4bE9X7%2BAQEJCvqzjSWa2Mg8MNNDA6eZ0I3fAeHv8QOFu8ydJUNrxnOTYAZOO57m%2BF%2FJ%2FQnIqJH1%2FUAjWFCdqUCVY4BekhI6S7nGC8tpARkjK8K4ujEmWoZjVZJBScqgPW5G0vuoOrpElWMQwtaWfccI26hSBU8mfEF5v9Mquo3q2UA9WgnwDE3IwRGDaAWNKCCtq2%2BMYpQI8jUsxb3amt9wYRRm7ZQKE7Gm6Xz4PR8vRjAR%2Ff4L0e1dpeYbpXh14sSE7Llkg%2BoxjpLbOThoPSvZZgoRKpwNGhw1maLGDseg9cA%2Fgsg3bprzl8qq9Rb0eogwdDZHa8gdEtojREx%2BiegONp%2FIfXzvGobKGqyHeGJYIlQwqB1iXqiuMeEXB9B8g8y3Ihnu%2Bk7cLmRf9asgCCzILJH47dqK6StFxKO4txwJliapO4hTmOY7rWHGI6OkfgWsv0M0vZt%2FkqEAvWifdOmJmTOOtL0ZPNUk09JyzdsnJnsbo%2BhVjkK4Ru%2B5XKKkRkm1lw4pBixG1jKIUMvxcp7E2y6oV5gTzk5VO5DfvXdRwjuLcapZJO01FQUOR11BUAGMpko5WHvt%2B3wtt3%2BMGHchy5MhyJMu%2BLH0tdzv9UMt5fcZ1yEoxcSrLoSwDWU7UxKZj80vL6m4LU7zOhE9zV0KUC8TVxjzkDFXHFieJmD6iKMcvcClVCSfcCZAkbMGoE0y4JIVLlI5g%2FG0tWWZMUkLlumAl%2F1o99%2BIdbdJMGVrVPmrRq41%2Buk7Pq1lc34B7PVIPIatVjl7FRyKb0mPGgTxHUvttWgyobPdRGGJOcswwETZcEsbI9jYLNe3PSCMokD1LcYbGZY5zkZFujxRllnUhuIKfGSj61wPF7UHBDAlGhDgTFM7CeZW43TdF3F6Tb%2F17iXvQUAR%2BLnG7dm4qmbugXleSsWOw9cAi476SeI5ia0Xhjh7J64UeoOWGxOPbDlPB3UvOBOFa1ORETyuMjBgyMzRPtZ5y9bI5NYJMYEwxQ41nbEOEHRjHKM%2Flw%2Bcbys5samaoHxlniixUJlJezBrUIXGMYYHe%2FjuNZ5pKHhHPHDd8ZxFt8HpEGl1jUp1eu2Zy3QONv4vZtmjMEcUcBuFRP5iBRzcS%2Be8M%2FHWI3Gkl8iGwiNNR%2FFol5efo%2FIZEXJRjNVERdlkp50o9SuGkprYaWXF5r518B%2BaS5nn6xiY024OhVn07ZVcn%2F3pgLRHhnVK0JqmHUPQg9Gte%2FoMUrfcDeqCutxv1QEPNQ2i8%2BzT%2FTr6%2BhAsw%2BTwHf2RD569h17NuDn9isDU3kmV0%2FTaJT%2FypkEgjX%2FnUtCy%2BS31cloLSqJ%2BKB0fnYd%2BkosZLo3xCGFwatnBp%2FwEvjVZkwf8V2TJs%2FCpk7e%2Bjt6UeOQ9D7FWedu0JSXAxA7k%2F2%2FDtbOPS3f6dbTyWMm%2F41vye3e%2BNuEvoNT8ODHrBfQ5jEZit6m6X4c3qB7ZiePUzJZj%2BCw%3D%3D)

## How to run?

Run this command to start and bind app on :8080 port. Database is a little bit slower at startup and app may restart several times before db is ready.

```bash
docker-compose up
```