# NotesMD Service

Fast  and lightweight backend for [NotesMD](https://github.com/RomanShirov/notesmd-app) app, writed on Fiber framework. Includes Prometheus and Grafana to collect metrics.

## Dependencies

+ Go (1.18 or newer)
+ Docker

## Stack

+ Fiber — API
+ Pgx — PostgreSQL
+ Goose — SQL Migrations
+ Bcrypt — Password hash generating
+ JWT — Authorization tokens

## Deploy

```sh
git clone https://github.com/RomanShirov/notesmd-service
cd notesmd-service
make frontend
```
* Add `.env` configuration file like `.env.example` and set your parameters.
* Go to `internal/web/notesmd-app/frontend` directory and create `.env` file with `VUE_APP_IP` field, where IP — IP Address of your backend server. Then:

```sh
make build
make run
```

This will install all dependencies for frontend, run the necessary Docker containers and run your application service.

For stop the application, use:
```sh
docker-compose down
```

## Additional Makefile commands

Clean frontend files:
```sh
make clear
```

Same as make clear, but also removing Docker containers (!) and drops db:
```sh
make reset
```

Run application with rebuild Docker containers (Required after modifying the Dockerfile)
```sh
make run-docker-build
```
