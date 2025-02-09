# container-monitor 
Данный проект предназначен для мониторинга Docker-контейнеров.
Состоит из двух компонентов:
- backend (текущий репозиторий) 
- pinger (https://github.com/DaniilStelmakh/pinger)

## backend 
Принимает GET,POST запросы по ендпоинту /pings:
GET /pings получает список всех записанных статусов по контейнерам
POST /pings записывает статус контейнера.

## pinger 
Принимает список ip и пингует их по протоколу icmp, полученный статус
отправляет на backend.

##  Запуск backend сервера
```shell
git clone https://github.com/DaniilStelmakh/container-monitor.git
```
Установить зависимости:
```shell
go mod download
```
Сборка приложения: 
```shell
go build -o /container-monitor/main.go
```
Запуск приложения:
```shell
/container-monitor/main
```
##  Запуск pinger сервиса 
```shell
git clone https://github.com/DaniilStelmakh/pinger
```
Установить зависимости:
```shell
go mod download
```
Сборка приложения:
```shell
go build -o /pinger/main.go
```
Запуск приложения:
```shell
/pinger/main
```
## Запуск через Docker 
```shell
docker compose up 
```
## Переменные окружения 
Backend:
- PORT порт на котором будут приниматься запросы
- DB_HOST хост базы данных
- DB_USER имя пользователя базы данных 
- DB_PASSWORD пароль пользователя базы данных
- DB_NAME схема

Pinger:
- INTERVAL время опроса контейнера
в секундах 
- IPS ip контейнеров через запятую 
- API_HOST хост backend сервиса
- API_PORT порт backend сервиса  