# How to run the program?

```
go run httpServer
```
Далее отправить запрос.
**Опции:**
1. через Postman, используя файл postman_collection.json из текущей директории
2. через консоль

Если запускать через консоль:

Создание события (create_event)
```
curl -X POST -H "Content-Type: application/json" -d '{"Id": "1", "Date": "2024-01-17T12:00:00Z"}' http://localhost:8080/create_event
```
Обновление события (update_event)

```
curl -X POST -H "Content-Type: application/json" -d '{"Id": "1", "Date": "2024-01-17T14:00:00Z"}' http://localhost:8080/update_event
```

Удаление события (delete_event)
```
curl -X POST -H "Content-Type: application/json" -d '{"Id": "1", "Date": "2024-01-17T14:00:00Z"}' http://localhost:8080/delete_event
```

Получение событий за день
```
curl http://localhost:8080/events_for_day?start_date=2024-01-17T00:00:00Z
```
Получение событий за неделю

```
curl http://localhost:8080/events_for_week?start_date=2024-01-17T00:00:00Z
```
Получение событий за месяц
```
curl http://localhost:8080/events_for_month?start_date=2024-01-17T00:00:00Z
```