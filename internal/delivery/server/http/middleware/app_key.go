package middleware

import (
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
	"net/http"
)

func ApiKeyMiddleware(config *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rContext := r.Context()
			keyConfig := config.ApiKey
			keyFromRequest := r.Header.Get("x-api-key")

			if keyFromRequest == "" {
				helper.ResponseHttp(w, helper.View{
					Status:       false,
					Code:         "401",
					Message:      "Invalid header key",
					ErrorMessage: nil,
					Data:         nil,
					Pagination:   helper.Pages{},
				}, http.StatusUnauthorized)
				return
			}

			if keyFromRequest != keyConfig {
				helper.ResponseHttp(w, helper.View{
					Status:       false,
					Code:         "401",
					Message:      "Invalid header key",
					ErrorMessage: nil,
					Data:         nil,
					Pagination:   helper.Pages{},
				}, http.StatusUnauthorized)

				return
			}
			next.ServeHTTP(w, r.WithContext(rContext))
		})
	}
}
