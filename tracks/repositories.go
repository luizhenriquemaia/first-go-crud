package tracks

import (
	"app/config"
	"database/sql"
)

type SqliteRepository struct {
	config.SqliteRepositoryAbstract
	db *sql.DB
}

func TrackRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) All() ([]Track, error) {
	rows, err := r.db.Query("SELECT * FROM track")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Track
	for rows.Next() {
		var track Track
		if err := rows.Scan(&track.ID, &track.Title, &track.Order_number, &track.Duration, &track.Album); err != nil {
			return nil, err
		}
		all = append(all, track)
	}
	return all, nil
}

func (r *SqliteRepository) Create(track Track) (*Track, error) {
	res, err := r.db.Exec(`
		INSERT INTO track(title, order_number, duration, album_id)
		SELECT ?, ?, ?, ?
		WHERE EXISTS(
			SELECT album.id FROM album WHERE album.id=?
		);
	`, track.Title, track.Order_number, track.Duration, track.Album, track.Album)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if affect == 0 {
		return nil, config.ErrCreateFailed
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	track.ID = id
	return &track, nil
}
