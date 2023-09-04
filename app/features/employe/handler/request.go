package handler 

type RequestCreateEmploye struct {
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
} 

type RequestUpdateEmploye struct { 
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
} 