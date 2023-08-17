package main

import (
	employeHandler "employe/app/features/employe/handler"
	employeRepo "employe/app/features/employe/repository"
	employeLogic "employe/app/features/employe/services"
	"employe/app/routes"
	"employe/config"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.GetConfiguration()
	db, _ := config.GetConnection(*cfg) 
	config.Migrate(db) 

	employeModel := employeRepo.New(db)
	employeServices := employeLogic.New(employeModel)
	employeController := employeHandler.New(employeServices, config.GetConfiguration())


	

	

	routes.Route(e, employeController, cfg) 

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
} 