package handler

import (
	annualreview "employe/app/features/annualReview"
	"employe/config"
	"employe/helper"
	middleware "employe/middlewares"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AnnualController struct {
	s annualreview.Service  
	config *config.Configuration 

}

func New(a annualreview.Service, c *config.Configuration) annualreview.Handler {
	return &AnnualController{s: a, config: c}
}

func (ar *AnnualController) CreateAnnual()echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestCreateAnnual 
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString) 
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		} 

		id := claims.ID

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input:", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		newAnnual := annualreview.Core{
			ID:         input.ID,
			EmplID:     input.EmplID,
			ReviewDate: input.ReviewDate,
		} 

		err = ar.s.CreateAnnual(newAnnual, id) 
		if err != nil {
			c.Logger().Error("Failed to create annual:", err)
			return c.JSON(http.StatusInternalServerError,helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 

		response := AnnualResponse{
			Code: http.StatusCreated, 
			Message: "Success created an annual",
			Data: AnnualData{
				ID: newAnnual.ID,
				EmplID : newAnnual.EmplID,
				ReviewDate : newAnnual.ReviewDate,
				
			},
		} 
		return c.JSON(http.StatusCreated, response)
	}
}

func (ar *AnnualController) GetAnnuals() echo.HandlerFunc {
	return func(c echo.Context) error {
		var annuals []annualreview.Core 
		var err error 

		// annuals, err = ar.s.GetAnnuals()
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		} 

		if len(annuals) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		formattedAnnuals := []ResponseGetAnnuals{}
		for _, annual :=  range annuals {
			formattedAnnual := ResponseGetAnnuals{
				ID        :annual.ID,
				EmplID    : annual.EmplID,
				ReviewDate:  annual.ReviewDate,
			}
			formattedAnnuals = append(formattedAnnuals, formattedAnnual )
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

		total := len(formattedAnnuals)
		totalPages := int(math.Ceil(float64(total)/ float64(perPageInt))) 

		startIndex := (pageInt - 1) * perPageInt 
		endIndex := startIndex + perPageInt 
		if endIndex > total {
			endIndex = total
		} 

		response := formattedAnnuals[startIndex:endIndex]

		pages := Pagination{
			Page: pageInt, 
			PerPage: perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, AnnualsResponse{
			Code: http.StatusOK,
			Message: "Successful operation.", 
			Data: response,
			Pagination: pages,
		})
	}
}

func (ar *AnnualController) GetEmplByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		emplID := claims.ID
		employee, err := ar.s.GetEmplByID(emplID)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		if employee == nil {
			c.Logger().Error("Employee not found")
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		annuals, err := ar.s.GetEmplByID(emplID) // Assuming you have a method to retrieve annuals for an employee
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error.", nil))
		}

		formattedAnnuals := make([]ResponseGetAnnuals, 0)
		for _, annual := range annuals {
			formattedAnnual := ResponseGetAnnuals{
				ID:         annual.ID,
				EmplID:     annual.EmplID,
				ReviewDate: annual.ReviewDate,
			}
			formattedAnnuals = append(formattedAnnuals, formattedAnnual)
		}

		page := c.QueryParam("page")
		perPage := c.QueryParam("per_page")
		if perPage == "" {
			perPage = "3"
		}
		pageInt, _ := strconv.Atoi(page)
		perPageInt, _ := strconv.Atoi(perPage)

		total := len(formattedAnnuals)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedAnnuals[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, AnnualsResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	}
}


func (ar *AnnualController) GetAnnual() echo.HandlerFunc {
	return func(c echo.Context) error {
		empl_ID, err := strconv.ParseUint(c.Param("id"), 10, 64) 
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}
        
		annualreview, err := ar.s.GetAnnual(uint(empl_ID))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))

		}

		response := ResponseGetAnnual{
			ID: annualreview.ID, 
			EmplID: annualreview.EmplID,
			ReviewDate: annualreview.ReviewDate,
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code: http.StatusOK, 
			Message: "Successful operation.",
			Data: response,
		})
	}
} 

func (ar *AnnualController)UpdateAnnual() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateAnnual 
		tokenString := c.Request().Header.Get("Authorization") 
		claims, err := middleware.ValidateJWT2(tokenString) 
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		id := claims.ID 
		if err  := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input from request body: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		updatedAnnual := annualreview.Core{
			ID: uint(id), 
			EmplID : input.EmplID,
			ReviewDate : input.ReviewDate, 

		}

		err =  ar.s.UpdateAnnual(updatedAnnual.ID, updatedAnnual)
		if err != nil {
			c.Logger().Error("Failed to updated emplID: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := RequestUpdateAnnual{
			ID: updatedAnnual.ID,
			EmplID: updatedAnnual.EmplID,
			ReviewDate: updatedAnnual.ReviewDate,
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code: http.StatusOK,
			Message: "Success updated an Annual Review",
			Data: response,
		})
	}
}

func (ar *AnnualController)DeleteAnnual() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middleware.ValidateJWT2(tokenString) 
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		} 

		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if claims.ID != uint(id) {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized,"Invalid Token",nil ))
		} 

		err = ar.s.DeleteAnnual(uint(id)) 
		if err != nil {
			c.Logger().Error("Error deleting annual", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internet Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an Annual", nil))
	}
}