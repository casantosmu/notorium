package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/casantosmu/notorium/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	logger *slog.Logger
	notes *models.NoteModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	pg := flag.String("pg", "postgres://notorium:notorium@localhost:5433/notorium", "Postgres connection string")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*pg)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }
	defer db.Close()

	app := &application{
		logger: logger,
		notes: &models.NoteModel{DB: db},
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if (err != nil) {
		return nil, err
	}

	err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

	return db, nil
}
