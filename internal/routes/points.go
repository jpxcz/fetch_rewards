package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/jpxcz/fetch_rewards/internal/models"
	"github.com/jpxcz/fetch_rewards/internal/repository/data"
	"github.com/julienschmidt/httprouter"
)

func pointsRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pointId := ps.ByName("id")
	v, err := data.GetId(pointId)
	if err != nil {
		slog.Error("error getting point data", slog.String("id", pointId), slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No receipt found for that id")
		return
	}

	response := models.RecieptPoints{
		Points: v,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
