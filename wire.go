package main

import (
	"github.com/goki0524/clean-arch/infrastructure"
	"github.com/goki0524/clean-arch/infrastructure/repository"
	"github.com/goki0524/clean-arch/presentation/controllers"
	"github.com/goki0524/clean-arch/usecases"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// DIの設定を行う
var set = wire.NewSet(
	/*
		Router
	*/
	infrastructure.NewRouter,

	/*
		Controller
	*/
	controllers.NewUsersController,

	/*
		Repository
	*/
	repository.NewCompanyBedRepository,
	wire.Bind(new(repository.ICompanyBedRepository), new(*repository.CompanyBedRepository)),

	/*
		UseCase
	*/
	usecases.NewUsersUseCase,
	wire.Bind(new(usecases.IUsersUseCase), new(*usecases.UsersUseCase)),
)

// Initialize DI
func Initialize(e *echo.Echo, db *gorm.DB) *infrastructure.Router {
	wire.Build(set)
	return &infrastructure.Router{}
}
