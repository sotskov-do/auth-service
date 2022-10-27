package models

import "github.com/go-chi/jwtauth/v5"

type TokenAuth struct {
	Token *jwtauth.JWTAuth
}
