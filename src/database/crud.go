package database

import (
	"database/sql"
	"fmt"
	"github.com/DaniilStelmakh/backend/src/dto"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// NewPing функция для добавления нового пинга в БД
func NewPing(db *sql.DB, ping dto.PingInfo) error {
	query := "INSERT INTO pings (ip_address, ping_time, last_seen) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, ping.Ip, ping.PingTime, ping.LastSeen)
	if err != nil {
		return fmt.Errorf("Ошибка при добавлении пинга: %s", err)
	}
	return nil
}

// CreateTable функция для создания таблицы в БД
func CreateTable(db *sql.DB, storagePath string) error {
	log.Printf("Storage path: %s", storagePath)

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS pings
		 (ip_address VARCHAR(50), ping_time FLOAT,
		  last_seen TIMESTAMP)`)
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func GetPings(db *sql.DB) ([]dto.PingInfo, error) {
	rows, err := db.Query("SELECT * FROM pings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pings []dto.PingInfo
	for rows.Next() {
		var ping dto.PingInfo
		err := rows.Scan(&ping.Ip, &ping.PingTime, &ping.LastSeen)
		if err != nil {
			log.Fatal("Ошибка при сканировании строки: %s", err)
			return nil, err
		}
		pings = append(pings, ping)
	}
	return pings, nil
}

// ConnectDB функция для подключения бд
func ConnectDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
