# Tasks API
A simple API written in Go with Gin Web Framework. The API use PostgreSQL as database.

## Deploy
Download the project:

    git clone https://github.com/nicolito128/tasks-api

Get the following packages:

    go get github.com/gin-gonic/gin
    go get github.com/lib/pq

Set a "DATABASE_URL" environment variable. See .env.example for more information on how to configure the database.

Start server:

    go run main.go

## Endpoints
You can play with some endpoints included here:

* GET: /tasks
* GET: /tasks/:id
* POST: /tasks
* PUT: /tasks/:id
* DELETE: /tasks/:id

## Interest links
* [gin-gonic/gin][1]

[1]: https://github.com/gin-gonic/gin