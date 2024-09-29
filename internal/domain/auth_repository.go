package domain

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/quartzeast/go-simple-banking/internal/pkg/log"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
	logger log.Logger
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)

	if resp, err := http.Get(u); err != nil {
		r.logger.Error("Error while sending...", "err", err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
			r.logger.Error("Error while decoding response from auth server", "err", err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

/*
  This will generate a url for token verification in the below format
  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}
*/

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{
		Host:   "localhost:8181",
		Path:   "/auth/verify",
		Scheme: "http",
	}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String()
}
