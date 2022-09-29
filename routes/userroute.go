package routes

import (
	"github.com/WytQuant/bunframework/controller"
	"github.com/WytQuant/bunframework/cookie"
	"github.com/WytQuant/bunframework/middlewares"
	"github.com/gorilla/sessions"
	"github.com/uptrace/bunrouter"
)

func UserRoute(app *bunrouter.Router) {

	cookie.Store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   3600 * 5,
	}

	userGroup := app.NewGroup("/user").Use(middlewares.NewCorsMiddleware)

	userGroup.POST("/register", controller.Register)
	userGroup.POST("/login", controller.Login)
	userGroup.POST("/logout", controller.Logout)
	userGroup.GET("/data", controller.GetUser)
}
