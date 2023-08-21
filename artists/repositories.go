package artists

import (
	"app/config"
	"database/sql"
	"errors"
)

var (
	ErrNotExists = errors.New("artista não existente")
)

type SqliteRepository struct {
	config.SqliteRepositoryAbstract
	db *sql.DB
}

func ArtistRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) All() ([]Artist, error) {
	rows, err := r.db.Query("SELECT * FROM artist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Artist
	for rows.Next() {
		var artist Artist
		if err := rows.Scan(&artist.ID, &artist.Name); err != nil {
			return nil, err
		}
		all = append(all, artist)
	}
	return all, nil
}

func (r *SqliteRepository) GetByID(id int64) (*Artist, error) {
	row := r.db.QueryRow("SELECT * FROM artist WHERE id=?", id)
	var artist Artist
	if err := row.Scan(&artist.ID, &artist.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &artist, nil
}

func (r *SqliteRepository) Create(artist Artist) (*Artist, error) {
	res, err := r.db.Exec(`
		INSERT INTO artist(id, name) 
		VALUES (?, ?)
	`, artist.ID, artist.Name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	artist.ID = id
	return &artist, nil
}
