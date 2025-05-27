package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/learies/go-shortener/internal/config"
	"github.com/learies/go-shortener/internal/dto"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewStorage(cfg *config.Config) (*PostgresStorage, error) {
	db, err := sql.Open("pgx", cfg.Postgres.URI)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("postgres ping failed: %w", err)
	}

	if err := initialize(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("postgres init failed: %w", err)
	}

	return &PostgresStorage{db: db}, nil
}

func initialize(db *sql.DB) error {
	createTableStmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS urls (
        id BIGSERIAL PRIMARY KEY,
        original_url TEXT NOT NULL,
        short_url VARCHAR(8) NOT NULL UNIQUE,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    )`)
	if err != nil {
		return err
	}
	defer createTableStmt.Close()

	if _, err = createTableStmt.Exec(); err != nil {
		return err
	}

	createIndexStmt, err := db.Prepare(`
    CREATE INDEX IF NOT EXISTS idx_original_url ON urls(original_url)`)
	if err != nil {
		return err
	}
	defer createIndexStmt.Close()

	if _, err = createIndexStmt.Exec(); err != nil {
		return err
	}

	createShortUrlIndexStmt, err := db.Prepare(`
    CREATE INDEX IF NOT EXISTS idx_short_url ON urls(short_url)`)
	if err != nil {
		return err
	}
	defer createShortUrlIndexStmt.Close()

	_, err = createShortUrlIndexStmt.Exec()
	return err
}

func (pg *PostgresStorage) Add(s dto.ShortenURLResponse) error {
	stmt, err := pg.db.Prepare(`
		INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.ShortURL, s.OriginalURL)
	return err
}
