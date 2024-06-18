package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServcer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServcer {
	return &ApiServcer{
		svc: svc,
	}
}
func (s *ApiServcer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}
func (s *ApiServcer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())

	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, fact.Fact)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)

}
