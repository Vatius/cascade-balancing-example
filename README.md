# Cascade balancing example
This is test work: GoLang, js

Start server (for example):

`./test-server.exe -addr=localhost:8081 -slave=http://localhost:8082 -max=2`

Здесь с помощью флагов указан андрес, на котором будет работать сервер, адрес сервера на который будут перенаправляться запросы и максимальное кол-во обрабатываемых запросов в секунду.

По заданию запустить в 3-х экземплярах, соответственно, три сервера с разными адресами и ссылками друг на друга, последний не имеет флага slave.

Клиент - надо переписать на го.

Request schema:
#### GET /
`[{
"price": int,
"quantity": int,
"amount": int,
"object": int,
"method:": int
}]`
