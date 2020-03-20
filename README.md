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

Смотреть [здесь](https://app.diagrams.net?lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1&title=Untitled%20Diagram.drawio#R7VhNc9sgEP01PsYjCcuyjpY%2FmkM7zYxn2uaIJWzTYOEiFNv59QUEEpKs2HWTtM00Bwae2AXeLm%2BJe2CyPXxgcLf5RBNEep6THHpg2vO8YRCKVgLHAgChVwBrhpMCcitggZ%2BQBh2N5jhBWW0ip5RwvKuDMU1TFPMaBhmj%2B%2Fq0FSX1VXdwjVrAIoakjX7FCd8U6Mh3KvwW4fXGrOw6%2BssWmskayDYwoXsLArMemDBKedHbHiaISO4ML4XdvONruTGGUn6JAYT8dpYvH8fRMppT%2F8tnyH7caC%2BPkOT6wIsN3QlkgdgjjpHeOj8aPhjN0wRJl04PRPsN5mixg7H8uhcJILAN3xIxckVXO0eMo0Pnrt2SC5FDiG4RZ0cxRRv4hk%2BdPwOTGPsqGq6heGNFYqgxqBNgXbquOBIdTdMvUDZoUTbOxU7%2BXsq84E9T5rcoa5EkbsdOdlcEHcby3gouUJro7jQmMMtwXOdK8MGO3ySvfd8M7%2B1v04MmvRgdzeiAuWUmRvfWl8pIDoxNsWGUtNSiERVxKJqzGJ2%2FcxyyNeLnEq0dZSuK%2FokgGowhAjl%2BrG%2F3VGT1CncUi4OUSTRo3rugkRzFMbWVLTtNR37DkddwVPDQcqQSrTz29bk3bOeeCGio2shRbaTakWoHBnd7o6HBRX8ufKhOYThT7Vi1vmqn2rCZ2OIe83raQoLXqcxpkTmICUDedixKzlh%2F2OIkkeYRQxl%2BgkvlSibhTpKkaPOjnj8VCIFLRCIYP6yVykwooUytC1bq7zlF0SVT%2B68KlZ3T3de5U35unL5Xi7hJ%2BGsz0kyhq1WGXiVHgrakx1wQ2SVS%2BZYUE6rYfZSBuKMZ5pjKGC4p53R7WYSa8ee0URRozglO0aR84zgvUynKV9YzxRW8ZaEYnS8UlxcFuyRYFaKjKFwu8GeFu4P0txFur6m3g2uFO2w4Am8r3G77baqUu5BeV4mxY6l12BLjkUY8R6u1lnDHzBT9wg8wuIV4YttDIrV7KZRguJY9ZegZh4FVQ%2BaW55nxU65eDmdWkfEtE7vUeNY2ZNmBcYyyTP3j84DSjk3NLfeRdaagxcpU4YVVWKfEsab5Zvvvv56NztUzxx3%2BYxUtfD0hDc4pqXleu%2Fbjug8af8%2B%2BtuXgDjEsaJAZ9dIv8OBCIf%2F%2FAn8RIXdOCvkYtITT0fpaPcq75PyCh7hsJ9pQC3bZKW2VH%2B1wWnNbzay0vH9afEN7Sfs8I2sTRu3B2Li%2BXLKrk3%2Ff8xMV4f1LdHhWosPhoJblvynRZt%2BgD%2Bp%2Bb4I%2BaLi5RsbFsPrpr5he%2FX4KZj8B)

## How to run?

Run this command to start and bind app on :8080 port. Database is a little bit slower at startup and app may restart several times before db is ready.

```bash
docker-compose up
```