// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/morio-pg/stubtool/src/application"
	"github.com/morio-pg/stubtool/src/application/usecase"
	"github.com/morio-pg/stubtool/src/domain"
	"github.com/morio-pg/stubtool/src/domain/service"
	"github.com/morio-pg/stubtool/src/infrastructure"
	"github.com/morio-pg/stubtool/src/infrastructure/firebase"
	"github.com/morio-pg/stubtool/src/interfaces"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/handler"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/middleware"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/router"
)

// Injectors from wire.go:

func InitializeRouter() router.Router {
	accountRepository := firebase.NewAccountRepository()
	authenticationUsecase := usecase.NewAuthenticationUsecase(accountRepository)
	authenticationMiddleware := middleware.NewAuthenticationMiddleware(authenticationUsecase)
	accountUsecase := usecase.NewAccountUsecase(accountRepository)
	accountHandler := handler.NewAccountHandler(accountUsecase)
	stubRepository := firebase.NewStubRepository()
	stubService := service.NewStubService(stubRepository)
	adminUsecase := usecase.NewAdminUsecase(stubRepository, stubService)
	adminHandler := handler.NewAdminHandler(adminUsecase)
	stubUsecase := usecase.NewStubUsecase(stubRepository)
	stubsHandler := handler.NewStubsHandler(stubUsecase)
	routerRouter := router.NewRouter(authenticationMiddleware, accountHandler, adminHandler, stubsHandler)
	return routerRouter
}

// wire.go:

var wireSet = wire.NewSet(application.WireSet, domain.WireSet, infrastructure.WireSet, interfaces.WireSet)
