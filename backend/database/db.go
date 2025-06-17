package database

import (
	"database/sql"
	"qr-backend/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

func NewDB() (*DB, error) {
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	db := &DB{conn: conn}
	if err := db.createTables(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS qr_codes (
		user_id TEXT PRIMARY KEY,
		random_string TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	)`

	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) UpsertQRCode(userID, randomString string) error {
	now := time.Now()
	query := `
	INSERT OR REPLACE INTO qr_codes (user_id, random_string, created_at, updated_at)
	VALUES (?, ?, 
		COALESCE((SELECT created_at FROM qr_codes WHERE user_id = ?), ?),
		?
	)`

	_, err := db.conn.Exec(query, userID, randomString, userID, now, now)
	return err
}

func (db *DB) GetQRCode(userID string) (*models.QRCode, error) {
	query := `SELECT user_id, random_string, created_at, updated_at FROM qr_codes WHERE user_id = ?`
	
	var qr models.QRCode
	err := db.conn.QueryRow(query, userID).Scan(
		&qr.UserID, &qr.RandomString, &qr.CreatedAt, &qr.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &qr, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}
