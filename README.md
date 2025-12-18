# crud is a pet project of basic crudl\* service to rule tasks written in go

> \*yeah there is also a List endpoint but `crud` sound just better, so live with that

## Main points:

- written using DI - Dependency Injection
- clean architecture with config, logger, storage, handlers as separeted modules
- initialization of all dependencies in `main.go`
- docker-compose with Postgres db service

# CRUDL Endpoints:

### **C - `/tasks` `POST`**

`Request:`

```json
{
  "name": "New note",
  "text": "today i wanted to eat the whole day"
}
```

`Response:`

```
{
    "id": 1,
    "name": "note",
    "text": "that's all!"
}
```

### **R - `/tasks/{i}` `GET`**

`Response:`

```json
{
  "name": "Final note",
  "text": "Today is my last note in this summer"
}
```

### **U - `/tasks/{i}` `PUT`**

`Request:`

```json
{
  "name": "Not final note",
  "text": "I changed my mind"
}
```

`Response:`
`Code:200`

### **D - `/tasks/{i}` `DELETE`**

`Request:`
`None`

`Response`:
`Code:200`

### **L - `/tasks` `GET`**

### **- `/tasks/?page={int}` `GET`**

`Request`None
`Response`:

```json
[
  { "Id": 1, "Name": "12.01.2025", "Text": "it's finally 2025" },
  { "Id": 2, "Name": "12.11.2025", "Text": "i don't belive i survived" }
]
```

### Errors:

`Response`:

```json
{
  "error": "error text"
}
```

## Fast start

- set up `.env` using `.env.example`
- install `docker` + `docker-compose`
- run `docker compose up --build` or add `-d` flag at the end to leave after all started
