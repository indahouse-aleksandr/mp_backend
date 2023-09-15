package repository

import (
	"context"
	"fmt"
	"os"
	"tc_backend/queries"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Q *queries.Queries
}

func NewRepository() *Repository {
	return &Repository{
		Q: CreateQueryBuilder(),
	}
}

func CreateQueryBuilder() *queries.Queries {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())

	return queries.New(conn)
}
