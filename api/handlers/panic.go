package handlers

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"

	cErrors "internauth-go/shared/errors"
)

// Panic handler catch panic to return a consistent response
func Panic(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Note: this defer should be first see "stack deferred"
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				cErrors.WriteHTTP(w, cErrors.ErrInternalError)
			}
		}()
		next(w, r, ps)
	}
}
