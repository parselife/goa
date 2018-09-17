package middleware

import (
	"github.com/kataras/iris/middleware/basicauth"
	"time"
)

var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
	Realm:   "Authorization Required", // defaults to "Authorization Required"
	Expires: time.Duration(30) * time.Minute,
})
