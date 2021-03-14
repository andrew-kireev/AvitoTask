# AvitoTask
# Описание проекта

REST API сервис работы со статистико

- Сервис написан по принципам clean architecture
- Настроено логирование, чтение конфига из .toml файла
- Уровни delivery, usecase протестированы с использованием моков, repository - с помощью sqlmock
- Для оптимизации использована распаковка json с помощью библиотеки easyjson
- Сервис запускается с помощью Docker и docker-compose

# Запуск

```
docker-compose up
```

# API


### Метод сохранения статистики

- Принимает на вход :
    - date - дату события
    - views - количество показов
    - clicks - количество кликов
    - cost - стоимость кликов (в рублях с точностью до копеек)

Поля views, clicks и cost - опциональные. Статистика агрегируется по дате.

Запрос:

```
curl --request POST \
  --url http://localhost:8080/api/v1/statistic \
  --header 'Content-Type: application/json' \
  --data '{
	"date": "2021-04-13",
	"views": "1100",
	"clicks": "10000",
	"cost": "10"
}'
```

Ответ:
- При успешном запросе: код 200, пустое тело ответа
- При ошибке валидации данных: код 400, пустое тело ответа
- При ошибке сервера: код 500, пустое тело ответа

### Метод показа статистики

- Параметры запроса:
    - from - дата начала периода (включительно)
    - to - дата окончания периода (включительно)
    - sort - опциональное поле, по которому будет происходить сортировка. 
      Возможны варианты:
      - sort_date - сортировка по дате
      - views - сортировка по показам
      - clicks - сортировка по кликам
      - cost - сортировка по стоимость кликов
      - cpc - сортировка по полю cpc
      - cpm - сортировка по полю cpm
    

- Параметры ответа:
    - date - дату события
    - views - количество показов
    - clicks - количество кликов
    - cost - стоимость кликов
    - cpc - cost/clicks (средняя стоимость клика)
    - cpm = cost/views * 1000 (средняя стоимость 1000 показов)
    

Запрос:

```
curl --request GET \
  --url 'http://localhost:8080/api/v1/statistic?from=2004-05-04&to=2005-05-04&sort=views' 
```


Ответ:
- При успешном запросе: код 200, в теле ответа json со статистикой
- При ошибке валидации данных: код 400, пустое тело ответа
- При ошибке сервера: код 500, пустое тело ответа
```
[
  {
    "date": "2005-03-13",
    "views": 1100,
    "clicks": 100000000,
    "cost": "0",
    "cpc": "0",
    "cpm": "0"
  },
  {
    "date": "2005-04-13",
    "views": 1100,
    "clicks": 100000000,
    "cost": "0",
    "cpc": "0",
    "cpm": "0"
  }
]
```

### Метод сброса статистики

- Удаляет всю сохраненную статистику


Запрос:

```
curl --request DELETE \
  --url http://localhost:8080/api/v1/statistic
```

Ответ:
- При успешном запросе: код 200, пустое тело ответа
- При ошибке сервера: код 500, пустое тело ответа
