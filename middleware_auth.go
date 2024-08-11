package main

import (
	"fmt"
	"net/http"

	"github.com/al4an2/goRssAggregator/auth"
	"github.com/al4an2/goRssAggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth apiKey is error: %s", err))
			return
		}

		user, err := apiCfg.DB.GerUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Getting user finish with error: %s", err))
			return
		}
		handler(w, r, user)
	}

}
