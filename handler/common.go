package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shandysiswandi/test-edufund-co-id/model"
)

func jsonResponse(w http.ResponseWriter, code int, v interface{}) {
	data, _ := json.Marshal(v)
	w.WriteHeader(code)
	w.Write(data)
}

func errorResponse(w http.ResponseWriter, err error) {
	errResp, ok := err.(model.ErrorResponse)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
		return
	}

	data, _ := json.Marshal(err)
	w.WriteHeader(errResp.Code)
	w.Write(data)
}

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
