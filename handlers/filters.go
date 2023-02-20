package handlers

import (
	"encoding/json"
	dto "finaltaskbe/dto/result"
	"finaltaskbe/repositories"
	"net/http"
)

type handlerFilter struct {
	FilterRepository repositories.FilterRepository
}

func HandlerFilter(FilterRepository repositories.FilterRepository) *handlerFilter {
	return &handlerFilter{FilterRepository}
}

func (h *handlerFilter) MultiFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()

	films, err := h.FilterRepository.MultiFilter(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: films}
	json.NewEncoder(w).Encode(response)
}
