package feedback

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

func feedbackNew(rw http.ResponseWrite, req *http.Request, ps httprouter.Params) {
  rw.ParseForm()

  model := {
    emoji: r.Form.Get("emoji"),
    message: r.Form.Get("message"),
    email: r.Form.Get("email")
  }

  go registerFeedback()

	rdr = render.New()
	rdr.JSON(rw, http.StatusCreated, {})
}


func registerFeedback() {

}
