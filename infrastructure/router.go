package infrastructure

import (
	"github.com/goki0524/clean-arch/presentation/controllers"
	"github.com/labstack/echo"
)

// Router ルーティング
type Router struct {
	echo            *echo.Echo
	usersController *controllers.UsersController
}

// NewRouter コンストラクタ
func NewRouter(echo *echo.Echo, usersController *controllers.UsersController) *Router {
	return &Router{echo: echo, usersController: usersController}
}

// Init ルーティング設定
func (r *Router) Init() {

	api := r.echo.Group("/api")
	{
		customer := api.Group("/customer")
		{
			customer.GET("/users", r.usersController.Get())
		}
	}
}
