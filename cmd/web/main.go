package main

import (
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
	"snippetbox.robertgleason.ca/internals/models"
)

// Add a snippets field to the application struct. This will allow us to
// use the SnippetModel type in our handlers.
type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// Load .env file if it exists (optional, but useful)
	err := godotenv.Load()
	if err != nil {
		// It’s okay if no .env file is present; just log it
		// You can comment this line if you don't want the log
		slog.Info("No .env file found, continuing...")
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "", "MySQL data source name (overrides environment)")

	flag.Parse()

	if *dsn == "" {
		*dsn = os.Getenv("MYSQL_DSN")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if *dsn == "" {
		logger.Error("No DSN provided. Set MYSQL_DSN or pass -dsn flag")
		os.Exit(1)
	}

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("Starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
