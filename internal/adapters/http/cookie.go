package http

import (
	"net/http"

	"github.com/sotskov-do/auth-service/internal/domain/auth/models"
)

func (s *Server) setCookie(w http.ResponseWriter, c models.Cookie) {
	cookie := http.Cookie{
		Name:     c.Name,
		Value:    c.Value,
		Path:     "/",
		HttpOnly: true,
		Expires:  c.Expiration,
	}

	http.SetCookie(w, &cookie)
}
