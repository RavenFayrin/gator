Connection String:      postgres://postgres:postgres@localhost:5432/gator
Goose Up:          goose postgres postgres://postgres:postgres@localhost:5432/gator up
    MUST RUN FROM:      gator/sql/schema
Goose Down:          goose postgres postgres://postgres:postgres@localhost:5432/gator down
    MUST RUN FROM:      gator/sql/schema
Generate SQL in Go:     sqlc generate
Run Commands:           go run . COMMAND AGRS
