package routes

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/receipts/process", processRoute)
	router.GET("/receipts/:id/points", pointsRoute)

	return router
}
