package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/sotskov-do/auth-service/internal/domain/auth/models"
)

func (s *Server) authHandlers() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", s.LoginHandler)
	r.Post("/register", s.RegisterHandler)

	return r
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	var l string

	cookie, err := r.Cookie("login_token")
	if err == nil {
		token, _ := jwtauth.VerifyToken(s.auth.GetToken().Token, cookie.Value)
		lRaw, ok := token.Get("login")
		if ok {
			l = lRaw.(string)
		}
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	user := &models.User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if l == user.Login {
		w.Write([]byte("\"status\": \"logged in by token\""))
		return
	}
	if !user.ValidateLogin() {
		http.Error(w, "\"error\": \"add login to request body\"", http.StatusBadRequest)
		return
	}
	if !user.ValidatePassword() {
		http.Error(w, "\"error\": \"add password to request body\"", http.StatusBadRequest)
		return
	}

	password, err := s.auth.Login(ctx, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, fmt.Errorf("\"error\": %s", err).Error(), http.StatusBadRequest)
		return
	}

	if password != user.EncodePassword() {
		log.Println("\"error\": \"wrong credentials\"")
		http.Error(w, "\"error\": \"wrong credentials\"", http.StatusForbidden)
		return
	}

	newToken, err := s.auth.GenToken(ctx, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "\"error\": \"server error\"", http.StatusInternalServerError)
		return
	}

	loginCookie := models.Cookie{
		Name:       "login_token",
		Value:      newToken,
		Expiration: time.Now().Add(time.Hour),
	}

	s.setCookie(w, loginCookie)

	w.Write([]byte("\"status\": \"logged in\""))
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, fmt.Errorf("\"error\": %s", err).Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	user := &models.User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, fmt.Errorf("\"error\": %s", err).Error(), http.StatusInternalServerError)
		return
	}

	if !user.ValidateLogin() {
		http.Error(w, "\"error\": \"add login to request body\"", http.StatusBadRequest)
		return
	}
	if !user.ValidateEmail() {
		http.Error(w, "\"error\": \"invalid email\"", http.StatusBadRequest)
		return
	}
	if !user.ValidatePassword() {
		http.Error(w, "\"error\": \"add password to request body\"", http.StatusBadRequest)
		return
	}
	if !user.ValidatePhone() {
		http.Error(w, "\"error\": \"invalid phone number\"", http.StatusBadRequest)
		return
	}

	err = s.auth.Register(ctx, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, fmt.Errorf("\"error\": %s", err).Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("\"status\": \"registration successfully completed\""))
}
