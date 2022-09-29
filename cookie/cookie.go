package cookie

import "github.com/gorilla/sessions"

var (
	Store          = sessions.NewCookieStore([]byte("super-secret-key"))
	AuthKey string = "authenticated"
	UserId  string = "user_id"
)
