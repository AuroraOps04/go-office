package main

import (
	"net/http"

	"github.com/AuroraOps04/go-office/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	// session
	store := cookie.NewStore([]byte("secret demo"))
	engine.Use(sessions.Sessions("office", store))
	userGroup := engine.Group("/user")
	{
		user := controller.NewUser()
		userGroup.POST("/save", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.GET("/current", user.CurrentUser)

	}
	_ = http.ListenAndServe(":4000", engine)
}
