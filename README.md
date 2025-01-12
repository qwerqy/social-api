# Social API built in Go

## File structure

`bin` is where compiled code will be.

`cmd` is where the executable will be. Main entry point, which is the `api` and
`migrate`. `api` will have everything server related. `migrate` is where all
migration related code will live.

`docs` is where swagger documentation will live.

`internal` is where internal code will be stored. These will not be
visible/exportable out of the project scope and will only be used within the
project scope.

`scripts` is where setting up the server script will be.

### Additional

`web` that contains frontend code. (React, Svelte, NextJS, etc)

## Principles

### Separation of concerns

Each level in your program should be separated by a clear barrier, the transport
layer, the service layer, the storage layer...

### Dependency Inversion Principle (DIP)

You're injecting the dependencies in your layers. You don't directly call them.
Why? It promotes loose coupling and makes it easier to test your programs.

### Adaptability to Change

By organizing your code in a modular and flexible way, you can more easily
introduce new features, refactor existing code, and respond to evolving business
requirements.

Your systems should be easy to change, if you have to change alot of existing
code to add a new feature you're doing it wrong.

### Focus on Business Value

And finally, focus on delivering value to your users, they are the ones who will
be paying your bills at the end of the month.. So focus on business value.

## Start local database

Run docker compose

## Migrations

We are using `migrate`

Example of creating a migration for create_users

```shell
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```

Migrating

```shell
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up
```

Clean up database version

```shell
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" force VERSION
# change the VERSION to the actual version you want to change
```

### Usage of `make`

Creates automated scripts for migrations. Refer to `Makefile`

To run migration

```shell
make migration create_users
```

To run migration up

```shell
make migrate-up
```

To run migration down

```shell
make migrate-down
```
