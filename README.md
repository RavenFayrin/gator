# Gator 
This program will allow the creation of users, feeds, and posts of RSS feeds. Check out below for installation instructions and a list of commands.

## Installation
This project requires **Postgres** and **Go** to run the program. You will also need access to a CLI tool. Use `go install` to instal gator.

Connection String:      postgres://postgres:postgres@localhost:5432/gator
Goose Up:          goose postgres postgres://postgres:postgres@localhost:5432/gator up
    MUST RUN FROM:      gator/sql/schema
Goose Down:          goose postgres postgres://postgres:postgres@localhost:5432/gator down
    MUST RUN FROM:      gator/sql/schema
Generate SQL in Go:     sqlc generate
Run Commands:           go run . COMMAND AGRS
