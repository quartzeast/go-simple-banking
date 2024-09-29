package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/quartzeast/go-simple-banking/internal/domain"
	"github.com/quartzeast/go-simple-banking/internal/pkg/apierr"
	"github.com/quartzeast/go-simple-banking/internal/response"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func NewAuthMiddleware(repo domain.AuthRepository) AuthMiddleware {
	return AuthMiddleware{repo}
}

func (a AuthMiddleware) AuthorizationHandler(routeName string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentRouteVars := vars(r)
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			token := getTokenFromHeader(authHeader)
			isAuthorized := a.repo.IsAuthorized(token, routeName, currentRouteVars)
			if isAuthorized {
				next.ServeHTTP(w, r)
			} else {
				response.Error(w, apierr.NewAPIError(apierr.CodeForbidden, errors.New("Unauthorized")))
			}
		} else {
			response.Error(w, apierr.NewAPIError(apierr.CodeForbidden, errors.New("missing token")))
		}
	})
}

func vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(0); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
