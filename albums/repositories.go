package albums

import (
	"app/config"
	"database/sql"
	"errors"
)

var (
	ErrNotExists = errors.New("Album não existente")
)

type SqliteRepository struct {
	config.SqliteRepositoryAbstract
	db *sql.DB
}

func AbumRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) All() ([]Album, error) {
	rows, err := r.db.Query("SELECT * FROM album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Artist, &album.Title, &album.Price); err != nil {
			return nil, err
		}
		all = append(all, album)
	}
	return all, nil
}

func (r *SqliteRepository) GetByID(id int64) (*Album, error) {
	row := r.db.QueryRow("SELECT * FROM album WHERE id=?", id)
	var album Album
	if err := row.Scan(&album.ID, &album.Artist, &album.Title, &album.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &album, nil
}

func (r *SqliteRepository) GetByArtist(artist string) ([]Album, error) {
	rows, err := r.db.Query("SELECT * FROM album WHERE artist=?", artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Artist, &album.Title, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (r *SqliteRepository) Create(album Album) (*Album, error) {
	res, err := r.db.Exec(`
			INSERT INTO album(title, artist, price)
			VALUES (?, ?, ?)
		`,
		album.Title, album.Artist, album.Price,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	album.ID = id
	return &album, nil
}

func (r *SqliteRepository) Update(id int64, dataUpdate Album) (*Album, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec(
		"UPDATE album SET title=?, artist=?, price=? WHERE id=?",
		dataUpdate.Title, dataUpdate.Artist, dataUpdate.Price, id,
	)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, config.ErrUpdateFailed
	}

	return &dataUpdate, nil
}

func (r *SqliteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM album WHERE id=?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotExists
		} else {
			return err
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return config.ErrDeleteFailed
	}

	return err
}
