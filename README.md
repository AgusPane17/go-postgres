
# Proyecto GO-POSTGRESQL-1

Este proyecto es una pequeña API REST en la cual se prueban distintas tecnologias y se hace un contacto en pos de aprender sobre las tecnologias utilizadas. 


## Stack de tecnologias

**Server:** Golang

**DB:** Docker, Postgresql
## Iniciar el proyecto

Una vez descargado el proyecto con git clone:

```bash
  git clone https://github.com/AgusPane17/go-postgres.git
```

Se necesita tener una base de datos en postgres y configurarla para salir por el puerto 8000

En mi caso baje una imagen de postgres de docker y cree un contenedor de postgres de la siguiente manera:

```bash
docker run --name myPostgres -e POSTGRES_USER=pane -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=panedb -p 8000:5432 -d postgres
```

Para iniciar el proyecto se utiliza el comando:

```bash
go run .
```
