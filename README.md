# NotesMD Service

Fast  and lightweight back-end for [NotesMD](https://github.com/RomanShirov/notesmd-app) app, writed on Fiber framework. Includes Prometheus and Grafana to collect metrics.

## Dependencies

+ Go (1.18 or newer)
+ Docker

## Stack

+ Fiber — API
+ Pgx — PostgreSQL
+ Goose — SQL Migrations
+ Bcrypt — Password hash generating
+ JWT — Authorization tokens

## Performance
![Tested in single thread. Average CPU load: 5-8%, RAM: 40-50 mb.](https://user-images.githubusercontent.com/47012273/203384982-e5f00ee8-dd34-48b4-bd9a-3b0f38639e62.png)

## Deploy

$ git clone https://github.com/RomanShirov/notesmd-service

$ make service-build

$ make make service-run

This will install all dependencies, download and build static front-end files, and run the necessary Docker containers.

## Additional Makefile commands

Clean `/build` and front-end files:

$ make clear

Same as make clear, but also removing Docker containers (!) and drops db:

$ make reset

Removes all build files, containers, rebuild and run a service:

$ make rebuild