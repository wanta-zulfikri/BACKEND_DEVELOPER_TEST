package handler

import (
	"employe/app/features/employe"
	"employe/config"
	"employe/helper"
	"employe/middlewares"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeController struct {
	s employe.Service
	config *config.Configuration
} 

func New(h employe.Service, c *config.Configuration) employe.Handler {
	return &EmployeController{s: h, config: c}
}


func (ec *EmployeController) CreateEmploye() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestCreateEmploye
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id := claims.ID

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 
		
		newEmploye := employe.Core{
			ID:            id,
			FirstName: input.FirstName,
			LastName: input.LastName,
			HireDate: input.HireDate,
			TerminationDate: input.TerminationDate,
			Salary: input.Salary,
	
		}

		err = ec.s.CreateEmploye(newEmploye, id)
		if err != nil {
			c.Logger().Error("Failed to create employe: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 

		response := EmployeResponse{
			Code:    http.StatusCreated,
			Message: "Success created an employe",
			Data: EmployeData{
				ID: newEmploye.ID,
				FirstName: newEmploye.FirstName,
				LastName: newEmploye.LastName,
				HireDate: newEmploye.HireDate,
				TerminationDate: newEmploye.TerminationDate,
				Salary: newEmploye.Salary,
			},
		}
		return c.JSON(http.StatusCreated, response)
	}
} 



func (ec *EmployeController) GetEmployes() echo.HandlerFunc {
	return func(c echo.Context) error {
		var employes []employe.Core
		var err error

		employes, err = ec.s.GetEmployes()
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}
		

		if len(employes) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		formattedEmployes := []ResponseGetEmployes{}
		for _, employe := range employes {
			formattedEmploye := ResponseGetEmployes{
				ID: employe.ID,
				FirstName: employe.FirstName,
				LastName: employe.LastName,
				HireDate: employe.HireDate,
				TerminationDate: employe.TerminationDate,
				Salary: employe.Salary,
				
			}
			formattedEmployes = append(formattedEmployes, formattedEmploye)
		}

		page := c.QueryParam("page")
		perPage := c.QueryParam("per_page")
		if page != "" || perPage == "" {
			perPage = "3"
		}
		pageInt := 1
		if page != "" {
			pageInt, _ = strconv.Atoi(page)
		}
		perPageInt, _ := strconv.Atoi(perPage)

		total := len(formattedEmployes)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedEmployes[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, EmployeeResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	}
} 


func (ec *EmployeController) GetEmployeByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		employeID := claims.ID
		employes, err := ec.s.GetEmployeByID(employeID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		if len(employes) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		formattedEmployes := []ResponseGetEmployes{}
		for _, employe := range employes {
			formattedEmploye := ResponseGetEmployes{
				ID: employe.ID,
				FirstName: employe.FirstName,
				LastName: employe.LastName,
				HireDate: employe.HireDate,
				TerminationDate: employe.TerminationDate,
				Salary: employe.Salary,
			}
			formattedEmployes = append(formattedEmployes, formattedEmploye)
		}

		page := c.QueryParam("page")
		perPage := c.QueryParam("per_page")
		if page != "" || perPage == "" {
			perPage = "3"
		}
		pageInt := 1
		if page != "" {
			pageInt, _ = strconv.Atoi(page)
		}
		perPageInt, _ := strconv.Atoi(perPage)

		total := len(formattedEmployes)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedEmployes[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, EmployeeResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	}
}

func (ec *EmployeController) GetEmploye() echo.HandlerFunc {
	return func(c echo.Context) error {
		employe_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		employe, err := ec.s.GetEmploye(uint(employe_id))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		response := ResponseGetEmploye{
			ID: employe.ID,
			FirstName: employe.FirstName,
			LastName: employe.LastName,
			HireDate: employe.HireDate,
			TerminationDate: employe.TerminationDate,
			Salary: employe.Salary,
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful operation.",
			Data:    response,
		})
	}
}

func (ec *EmployeController) UpdateEmploye() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateEmploye
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		id := claims.ID
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input from request body: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		updatedEmploye := employe.Core{
			ID:            uint(id), 
			FirstName:         input.FirstName,
			LastName:          input.LastName,
			HireDate:          input.HireDate,
			TerminationDate:   input.TerminationDate,
			Salary:            input.Salary,
			
		}

		err = ec.s.UpdateEmploye(updatedEmploye.ID, updatedEmploye)
		if err != nil {
			c.Logger().Error("Failed to update employe: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := ResponseUpdateEmployes{
			ID: updatedEmploye.ID,
			FirstName: updatedEmploye.FirstName,
			LastName: updatedEmploye.LastName,
			HireDate: updatedEmploye.HireDate,
			TerminationDate: updatedEmploye.TerminationDate,
			Salary: updatedEmploye.Salary,
		}
		

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Success updated an employe.",
			Data:    response,
		})
	}
}

func (ec *EmployeController) DeleteEmploye() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if claims.ID != uint(id) {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is not valid for this user.", nil))
		}

		err = ec.s.DeleteEmploye(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting employe", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an employe", nil))
	}
}