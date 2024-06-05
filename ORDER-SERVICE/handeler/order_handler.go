package handeler

import (
	"encoding/json"
	"log"
	"net/http"
	"taqsir/model"
)

func Order(w http.ResponseWriter, r *http.Request) {
	var cation model.Cation
	json.NewDecoder(r.Body).Decode(&cation)

	log.Printf("Xabarnoma qabul qilindi: User %s %s", cation.UserID, cation.Action)
	w.WriteHeader(http.StatusOK)
}