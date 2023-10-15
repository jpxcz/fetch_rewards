package routes

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/jpxcz/fetch_rewards/internal/models"
	"github.com/jpxcz/fetch_rewards/internal/repository/data"
	"github.com/jpxcz/fetch_rewards/internal/repository/reciept"
	"github.com/julienschmidt/httprouter"
)

func processRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// HandleError
		return
	}

	var reciept reciept.Reciept
	err = json.Unmarshal(body, &reciept)
	if err != nil {
		slog.Error("error on unmarshal request", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The reciept is invalid")
		return
	}

	p, err := reciept.CreatePoints()
	if err != nil {
		slog.Error("error on creating string points", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The reciept is invalid")
		return
	}

	id := data.GetNewId()
	_ = data.SetId(id, p)

	response := models.RecieptInsertResponse{
		Id: id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
