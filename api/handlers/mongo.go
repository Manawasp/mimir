package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"mimir/api/config"
	"mimir/api/database"
)

// Mongo handler copy (take from the pool see https://godoc.org/labix.org/v2/mgo)
// the session and insert it inside of the context
func Mongo(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session := database.Session.Copy()
		defer session.Close()

		newCtx := database.ToContext(r.Context(), session.DB(config.App.DB.Database))
		next(w, r.WithContext(newCtx), ps)
	}
}
