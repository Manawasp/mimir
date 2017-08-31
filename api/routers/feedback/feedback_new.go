package feedback

import (
	"encoding/json"
	"feedback/api/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func feedbackNew(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var feed models.Feedback

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&feed); err != nil {
		return
	}
	defer req.Body.Close()

	go registerFeedback()

	rw.WriteHeader(http.StatusCreated)
}

func registerFeedback() {

}
