# UrlCutterAPI

Тестовое задание для Ozon.

Создает API сервер и базу данных postgres.
Сервер принимает методы GET и POST. Ссылки передаются в теле запроса:

curl -X POST -d "https://www.youtube.com/watch?v=et4esfqc43o" localhost:8000

curl -X GET -d "https://ozon.cc/givemeajob" localhost:8000

Длина каталога согласно ТЗ - 10 символов.

Инструкция по использованию:

$ make

Для докера:

$ ./build - первый запуск и сборка

Все следующие запуски:

$ docker-compose run [options] --service-ports api

example:

$ docker-compose run -e SaveMethod=inMemory --service-ports api


default of  env. variables:
  	- SaveMethod=postgres
  	- DBHost=db
  	- DBPort=5432
  	- ServerPort=8000
  	- DBName=postgres


