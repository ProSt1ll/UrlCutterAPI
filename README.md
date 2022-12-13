# UrlCutterAPI

Тестовое задание для Ozon.

Создает API сервер и базу данных postgres.
Сервер принимает методы GET и POST. Ссылки передаются в теле запроса:

curl -X POST -d "https://www.youtube.com/watch?v=et4esfqc43o" localhost:8000

curl -X GET -d "https://ozon.cc/givemeajob" localhost:8000

Длина каталога согласно ТЗ - 10 символов.

Инструкция по использованию:

$ make

$ ./urlcutterapi

$ docker-compose up --build

$ docker-compose run [options] urlcutterapi

example:

$ docker-compose run -e SaveMethod=inMemory urlcutterapi

default of  env. variables:
  	- SaveMethod=postgres
  	- DBHost=db
  	- DBPort=5432
  	- ServerPort=8000
  	- DBName=postgres


