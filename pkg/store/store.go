package store

import (
	"database/sql"

	"github.com/dzsak/url-shortener/pkg/model"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	*sql.DB
}

func New(dataSourceName string) (*Store, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	store := &Store{db}
	err = store.createUrlsTable()
	return store, err
}

func NewTest() (*Store, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	store := &Store{db}
	err = store.createUrlsTable()
	return store, err
}

func (store *Store) createUrlsTable() error {
	createUrlsTableSQL := `CREATE TABLE IF NOT EXISTS urls (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"original" TEXT,
		"shortKey" TEXT
	  );`

	statement, err := store.Prepare(createUrlsTableSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

func (store *Store) InsertUrl(url model.Url) error {
	insertStudentSQL := `INSERT INTO urls(original, shortKey) VALUES (?, ?)`
	statement, err := store.Prepare(insertStudentSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(url.Original, url.ShortKey)
	return err
}

func (store *Store) GetUrlByOriginal(original string) (model.Url, error) {
	getUrlSQL := `SELECT original, shortKey FROM urls WHERE original = ?`
	row := store.QueryRow(getUrlSQL, original)

	var url model.Url
	err := row.Scan(&url.Original, &url.ShortKey)
	return url, err
}

func (store *Store) GetUrlByShortKey(shortKey string) (model.Url, error) {
	getUrlSQL := `SELECT original, shortKey FROM urls WHERE shortKey = ?`
	row := store.QueryRow(getUrlSQL, shortKey)

	var url model.Url
	err := row.Scan(&url.Original, &url.ShortKey)
	return url, err
}
