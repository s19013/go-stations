package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// これは必要なかった｡handlerディレクトリのhealthz.goのServeHTTPに処理書けって書いてた
// 説明読んでなかった私が悪い｡記念に残そ｡
func Healthz(writer http.ResponseWriter, req *http.Request) {

	// レスポンスを生成する
	response := model.HealthzResponse{Message: "OK"}

	// データ構造をJSONにエンコードしてHTTPレスポンスとして返す
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Println(err)
		return
	}

}
