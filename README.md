# Хакатон

![\[img\](https://moe.video/storage/WKwwWl4HXH4vDd8iENiR.svg)](https://moe.video/storage/WKwwWl4HXH4vDd8iENiR.svg)

<details>

<summary>Ссылки по README</summary>

>
> [**Установка и запуск**](#установка-и-запуск)
>
> [**Инструкция пользования**](#инструкция-пользования)
>
> [**Структура приложения**](#структура-приложения)
</details>

## Установка и запуск

### С использованием Docker

1. Склонируйте репозиторий
2. Запустите Docker
3. Выполните

    ```shell
    docker compose up
    ```

    **Готово!**

### Без использования Docker

1. Склонируйте репозиторий
2. Заполните переменные окружения по примеру из `.env.example`
3. Запустите сервер PostgreSQL
4. Выполните

    ```shell
    go mod download
    go run .\src\main.go
    ```

     **Готово!**


*Также, если требуется заменить категории, достаточно заменить содержимое файла categories.json не меняя его структуру*

## Инструкция пользования

> [**Вернуть в начало**](#хакатон)

Заполнение базы данных категориями в первый раз осуществляется запуском программы с аргументом seed.

```shell
go run src\main.go seed

.\main.exe seed
```


### Эндпоинты

`GET` /categories

#### Query параметры

- url

    > Обязательный
    >
    > URL сайта, категории которого нужно получить

- confident

    > Необязательный
    >
    > Если выбрать false, то в ответе будут категории, в которых мы не уверены

#### Формат ответа

Ответ приходит в виде JSON со списком ID категорий под ключом "categories"

`GET` /categories/?url=https://playground.ru

```json
{
    "categories": [
        1972
    ]
}
```

`GET` /categories/?url=https://playground.ru&confident=false

```json
{
  "categories": [
    1742,
    1922,
    1972,
    2832,
    1722
  ]
}
```

#### Ошибки

`GET` /categories/

> Параметр URL не задан

```json
{
  "error": "URL is required"
}
```

`GET` /categories/?url=https:/asdf

> URL некорректен

```json
{
  "error": "Invalid URL"
}
```

`GET` /categories/?url=https://kniga-online.com

> По URL не проходят запросы

```json
{
  "error": "URL is unreachable"
}
```

## Структура приложения

> [**Вернуть в начало**](#хакатон)

### Файловая структура

```text
src/
├─ server/
│  ├─ server.go - сервер Gin
├─ config/
│  ├─ config.go - файл собирающий конфигурацию из окружения
├─ database/
│  ├─ database.go - модуль для работы с базой данных PostgreSQL, использующий gorm
├─ categorizes/
│  ├─ skydns/
│  │  ├─ skydns.go - модуль для работы с API skydns для категоризации сайта
├─ main.go - файл для запуска программы
```