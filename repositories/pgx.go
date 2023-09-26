package repositories

import (
	"database/sql"
	"os"

	"github.com/indahouse-aleksandr/mp_backend/queries"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Pgx struct {
	queries.DBTX
	Q *queries.Queries
}

func NewPgx() *Pgx {
	if pool, err := sql.Open("pgx/v5", os.Getenv("DATABASE_URL")); err == nil {
		return &Pgx{
			Q: queries.New(pool),
		}
	}

	return nil
}
