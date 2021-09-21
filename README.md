# Cascade balancing example

This is test GoLang work

### Server

`# make`

`# bin/server -addr=localhost:8081 -slave=http://localhost:8082 -max=2`

Здесь с помощью флагов указан андрес, на котором будет работать сервер, адрес сервера на который будут перенаправляться запросы и максимальное кол-во обрабатываемых запросов в секунду.

Можно запускать несколько экземпляров с разными адресами и ссылками друг на друга, последний не имеет флага slave.

### Client

Client send test data to server

`# bin/client -interval=1s -url=http://localhost:8080/`

### Request schema:

#### POST /
`[{
"price": int,
"quantity": int,
"amount": int,
"object": int,
"method:": int
}]`
