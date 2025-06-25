package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"qr-backend/models"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func NewDB() (*DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf(".envファイルの読み込みに失敗しました（環境変数を使用します）: %v", err)
	}

	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbPort := getEnvOrDefault("DB_PORT", "5432")
	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "password")
	dbName := getEnvOrDefault("DB_NAME", "qr_code_db")
	dbSSLMode := getEnvOrDefault("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("データベース接続に失敗しました: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("データベースへの接続確認に失敗しました: %w", err)
	}

	db := &DB{conn: conn}
	if err := db.createTables(); err != nil {
		return nil, err
	}

	return db, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (db *DB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS qr_codes (
		user_id VARCHAR(255) PRIMARY KEY,
		random_string VARCHAR(10) NOT NULL,
		status INTEGER NOT NULL DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.conn.Exec(query)
	if err != nil {
		return fmt.Errorf("テーブル作成に失敗しました: %w", err)
	}

	indexQuery := `
	CREATE INDEX IF NOT EXISTS idx_qr_codes_user_id_random_string 
	ON qr_codes(user_id, random_string)`

	_, err = db.conn.Exec(indexQuery)
	if err != nil {
		return fmt.Errorf("インデックス作成に失敗しました: %w", err)
	}

	return nil
}

func (db *DB) UpsertQRCode(userID, randomString string) error {
	now := time.Now()
	query := `
	INSERT INTO qr_codes (user_id, random_string, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (user_id) 
	DO UPDATE SET 
		random_string = EXCLUDED.random_string,
		status = 0,
		updated_at = EXCLUDED.updated_at`

	_, err := db.conn.Exec(query, userID, randomString, 0, now, now)
	if err != nil {
		return fmt.Errorf("QRコードデータの保存に失敗しました: %w", err)
	}
	return nil
}

func (db *DB) GetQRCode(userID string) (*models.QRCode, error) {
	query := `SELECT user_id, random_string, status, created_at, updated_at FROM qr_codes WHERE user_id = $1`
	
	var qr models.QRCode
	err := db.conn.QueryRow(query, userID).Scan(
		&qr.UserID, &qr.RandomString, &qr.Status, &qr.CreatedAt, &qr.UpdatedAt,
	)
	
	if err != nil {
		return nil, fmt.Errorf("QRコードデータの取得に失敗しました: %w", err)
	}
	
	return &qr, nil
}

func (db *DB) VerifyAndUpdateQRCode(userID, randomString string) (bool, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return false, fmt.Errorf("トランザクション開始に失敗しました: %w", err)
	}
	defer tx.Rollback()

	var count int
	query := `SELECT COUNT(*) FROM qr_codes WHERE user_id = $1 AND random_string = $2 AND status = 0`
	err = tx.QueryRow(query, userID, randomString).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("QRコード検証に失敗しました: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	updateQuery := `UPDATE qr_codes SET status = 1, updated_at = CURRENT_TIMESTAMP WHERE user_id = $1 AND random_string = $2 AND status = 0`
	_, err = tx.Exec(updateQuery, userID, randomString)
	if err != nil {
		return false, fmt.Errorf("QRコードステータスの更新に失敗しました: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return false, fmt.Errorf("トランザクションのコミットに失敗しました: %w", err)
	}

	return true, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}
