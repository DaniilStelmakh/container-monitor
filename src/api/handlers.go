package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DaniilStelmakh/container-monitor/src/database"
	_ "github.com/DaniilStelmakh/container-monitor/src/database"
	"github.com/DaniilStelmakh/container-monitor/src/dto"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// ShowPingsHandler хендлер для получения списка пингов
func ShowPingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pings, err := database.GetPings(db)
		if err != nil {
			log.Printf("Ошибка при получении пингов: %s", err)
			return
		}
		jsn, err := json.Marshal(&pings)
		if err != nil {
			log.Printf("Ошибка при маршалинге пингов: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)
	}
}

// WritePingHandler хендлер для записи нового пинга
func WritePingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var pingDate dto.PingInfo
		var buf bytes.Buffer
		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			log.Printf("Ошибка при чтении тела запроса: %s", err)
			return
		}
		if err = json.Unmarshal(buf.Bytes(), &pingDate); err != nil {
			log.Printf("Ошибка при размаршалинге тела запроса: %s", err)
			return
		}

		err = database.NewPing(db, pingDate)
		if err != nil {
			log.Printf("Ошибка при добавлении пинга: %s", err)
			return
		}

	}

}
