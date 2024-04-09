# effectiveMobile-test

Реализовать каталог автомобилей. Необходимо реализовать следующее
1. Выставить rest методы
2. Получение данных с фильтрацией по всем полям и пагинацией
3. Удаления по идентификатору
4. Изменение одного или нескольких полей по идентификатору
5. Добавления новых автомобилей в формате

```json
{
"regNums": ["X123XX150"] // массив гос. номеров
}
```
6. При добавлении сделать запрос в АПИ, описанного сваггером

```yaml
openapi: 3.0.3
info:
title: Car info
version: 0.0.1
paths:
/info:
get:
parameters:
- name: regNum
in: query
required: true
schema:
type: string
responses:
'200':
description: Ok
content:
application/json:
schema:
$ref: '#/components/schemas/Car'
'400':
description: Bad request
'500':
description: Internal server error
components:
schemas:
Car:
required:
- regNum
- mark
- model
- owner
type: object
properties:
regNum:
type: string
example: X123XX150
mark:
type: string
example: Lada
model:
type: string
example: Vesta
year:
type: integer
example: 2002
owner:
$ref: '#/components/schemas/People'
People:
required:
- name
- surname
type: object
properties:
name:
type: string
surname:
type: string
patronymic:
type: string
```
7. Обогащенную информацию положить в БД postgres (структура БД должна быть
создана путем миграций при старте сервиса)
8. Покрыть код debug- и info-логами
9. Вынести конфигурационные данные в .env-файл
10. Сгенерировать сваггер на реализованное АПИ
