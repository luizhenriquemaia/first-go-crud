package config

import (
	"database/sql"
	"errors"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("record doesn't exists")
	ErrCreateFailed = errors.New("create failed")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SqliteRepositoryAbstract interface {
	All()
	GetByID()
	Create()
	Update()
	Delete()
}

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS artist(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(50) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS album(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(50) NOT NULL,
			price DECIMAL(10,2) DEFAULT '0.00',
			artist_id INT, 
			FOREIGN KEY (artist_id) REFERENCES artist(id)
		);
	`
	_, err := r.db.Exec(query)
	return err
}
