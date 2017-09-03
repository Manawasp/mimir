package feedback

import (
	"encoding/json"
	"feedback/api/config"
	"feedback/api/database"
	"feedback/api/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func feedbackNew(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	feed := &models.Feedback{
		URL:    req.URL.String(),
		Domain: req.URL.Host,
	}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(feed); err != nil {
		return
	}
	defer req.Body.Close()

	go registerFeedback(feed)

	rw.WriteHeader(http.StatusCreated)
}

func registerFeedback(feed *models.Feedback) {
	if database.Session != nil {
		s := database.Session.Copy()
		db := s.DB(config.App.DB.Database)
		defer s.Close()

		feed.Insert(db)
	}

	for _, hook := range config.App.Hooks {
		hook.Notify(feed)
	}
}
