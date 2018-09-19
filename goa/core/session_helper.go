package core

import "github.com/kataras/iris/sessions"

const (
	UserId        = "UserID"
	AUTHENTICATED = "Authenticated"
	IsAdmin       = "IsAdmin"
)

func GetCurrentUserID(s *sessions.Session) int64 {
	return s.GetInt64Default(UserId, 0)
}

func IsLoggedIn(s *sessions.Session) bool {
	return s.GetInt64Default(UserId, 0) > 0
}

func Admin(s *sessions.Session) bool {
	return s.GetBooleanDefault(IsAdmin, false)
}
