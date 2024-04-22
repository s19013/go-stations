package controller

import (
	"encoding/json"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

func Healthz(writer http.ResponseWriter, req *http.Request) {
	// log.Println("writer", writer)

	// レスポンスを生成する
	response := model.HealthzResponse{Message: "message"}

	// 2. データ構造をJSONにエンコードしてHTTPレスポンスとして返す
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		return
	}

	// なんかこんな書き方もあるみたい｡
	// jsonData, err := json.Marshal(person)
	// if err != nil {
	// 	fmt.Println("エラー:", err)
	// 	return
	// }
}
