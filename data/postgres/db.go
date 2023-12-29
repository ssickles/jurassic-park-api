package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	"log"
	"os"
)

type PgDb interface {
	Model(model ...interface{}) *pg.Query
	Exec(query interface{}, params ...interface{}) (pg.Result, error)
	RunInTransaction(ctx context.Context, fn func(*pg.Tx) error) error
}

/*
CreateDbConnection creates a new database connection to the postgres database.

	If POSTGRES_URL is not set, it will default to a local database.
	If POSTGRES_DEBUG is set to true, it will log all failed queries to stderr.
*/
func CreateDbConnection() (*pg.DB, error) {
	purl := os.Getenv("POSTGRES_URL")
	if purl == "" {
		purl = "postgres://jurassic-park:jp@0.0.0.0:5442/jurassic-park?sslmode=disable"
	}
	opt, err := pg.ParseURL(purl)
	if err != nil {
		return nil, fmt.Errorf("invalid postgres URL: %w", err)
	}
	db := pg.Connect(opt)
	if os.Getenv("POSTGRES_DEBUG") == "true" {
		db.AddQueryHook(pgdebug.DebugHook{
			Verbose: true,
		})
	} else {
		db.AddQueryHook(DebugHook{FailedQueryLogger: log.New(os.Stderr, "[FAILED_QUERY] ", 0)})
	}
	return db, nil
}

var _ pg.QueryHook = (*DebugHook)(nil)

/*
DebugHook is a helper struct for logging failed queries.

	It follows the pg.QueryHook interface so that it can be added as a hook for a go-pg connection.
*/
type DebugHook struct {
	FailedQueryLogger *log.Logger
}

func (d DebugHook) BeforeQuery(ctx context.Context, event *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DebugHook) AfterQuery(ctx context.Context, event *pg.QueryEvent) error {
	if d.FailedQueryLogger != nil {
		if event.Err != nil {
			q, err := event.FormattedQuery()
			if err != nil {
				return err
			}
			d.FailedQueryLogger.Println(string(q))
		}
	}
	return nil
}
