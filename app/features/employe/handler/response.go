package handler

type ResponseGetEmployes struct {
	ID              uint     `json:"id"` 
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
} 

type ResponseGetEmploye struct {
	ID              uint     `json:"id"` 
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
} 


type EmployeResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    EmployeData  `json:"data"`
}

type EmployeData struct {
	ID              uint     `json:"id"` 
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
}



type EmployeeResponse struct {
	Code       int                 `json:"code"`
	Message    string              `json:"message"`
	Data       []ResponseGetEmployes `json:"data"`
	Pagination Pagination          `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type ResponseUpdateEmployes struct {
	ID              uint     `json:"id"` 
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
}