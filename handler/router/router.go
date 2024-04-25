package router

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		todoService := service.NewTODOService(todoDB)
		todoHandler := handler.NewTODOHandler(todoService)

		if r.Method == http.MethodPost {
			// body取り出し
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}

			// json変化
			var request model.CreateTODORequest
			err2 := json.Unmarshal(body, &request)
			// err2 := json.NewDecoder(bytes.NewReader(body)).Decode(&request)
			if err2 != nil {
				log.Println("err2", err2)
				http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
				return
			}

			// subjectがあるか確認
			if request.Subject == "" {
				http.Error(w, "Error subject not exist", http.StatusBadRequest)
				return
			}

			// 登録してtodoを受け取る
			response, err3 := todoHandler.Create(r.Context(), &request)
			if err3 != nil {
				log.Println("err3", err3)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			err4 := json.NewEncoder(w).Encode(response)
			if err4 != nil {
				log.Println("err4", err4)
				return
			}

		}
	})

	return mux
}
