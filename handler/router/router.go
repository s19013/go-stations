package router

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/model"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// body取り出し
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}

			// json変化
			var request model.CreateTODORequest
			err2 := json.Unmarshal([]byte(body), &request)
			if err2 != nil {
				http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
				return
			}

			// subjectがあるか確認
			if request.Subject == "" {
				http.Error(w, "Error subject not exist", http.StatusBadRequest)
				return
			}

			// 登録してtodoを受け取る

		}
	})

	return mux
}
