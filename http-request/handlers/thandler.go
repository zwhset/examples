package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"qianbao.com/examples/http-request/defs"
)

func TestHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var u defs.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		defs.ErrHttpResponse(w, defs.ErrJSONParseFailed)
		return
	}

	if u.Username != "zwhset" || u.Password != "xxx" {
		defs.ErrHttpResponse(w, defs.ErrNotAuthUser)
		return
	}

	defs.HttpResponse(w, defs.ResJSON)
	return
}
