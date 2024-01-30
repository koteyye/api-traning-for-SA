# api-traning-for-SA
Локальный сервер-тренажер для системных аналитиков

## Запуск через Docker
```
docker run -it -d -e SERVER_ADDRESS="0.0.0.0:8081" -p 8081:8081  koteyye/apitraning
```

##
По умолчанию контейнер виден на порту 8081, при необходимости можно изменить в коменде

## Swagger
http://localhost:8081/swagger/index.html