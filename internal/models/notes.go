package models

import (
	"database/sql"
	"errors"
	"time"
)

type Note struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type NoteModel struct {
	DB *sql.DB
}

func (m *NoteModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO notes (title, content, created_at, expires_at)
	VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + $3 * INTERVAL '1 day')
	RETURNING note_id`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *NoteModel) Get(id int) (Note, error) {
	stmt := `SELECT note_id, title, content, created_at, expires_at FROM notes
    WHERE expires_at > CURRENT_TIMESTAMP AND note_id = $1`

	var n Note
	err := m.DB.QueryRow(stmt, id).Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Note{}, ErrNoRecord
		}
		return Note{}, err
	}

	return n, nil
}

func (m *NoteModel) Latest() ([]Note, error) {
	stmt := `SELECT note_id, title, content, created_at, expires_at FROM notes
    WHERE expires_at > CURRENT_TIMESTAMP ORDER BY note_id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note

	for rows.Next() {
		var n Note
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.ExpiresAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
