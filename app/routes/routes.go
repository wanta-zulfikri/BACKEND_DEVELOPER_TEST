package routes

import (
	"employe/app/features/employe"
	"employe/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
) 


func Route(e *echo.Echo, ec employe.Handler, config *config.Configuration) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger()) 
	 

	//employe 
	e.GET("/employe", ec.GetEmployes()) 
	e.POST("/employe", ec.CreateEmploye()) 
	e.GET("/employe/:id", ec.GetEmploye()) 
	e.PUT("/employe/:id", ec.UpdateEmploye())
	e.DELETE("/employe/:id", ec.DeleteEmploye())
}