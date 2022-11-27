package routes

import (
	"net/http"

	"github.com/assyatier21/user-deall-technical-test/controller"
	"github.com/assyatier21/user-deall-technical-test/controller/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes() *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	e.POST("v1/login", auth.Login)
	e.POST("v1/register", auth.Register)

	e.GET("v1/article/get", controller.GetArticleByID)
	e.GET("v1/article/getUserPoints", controller.GetPointsByUserId)
	e.POST("v1/article/create", controller.InsertArticle)
	return e
}

func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}
