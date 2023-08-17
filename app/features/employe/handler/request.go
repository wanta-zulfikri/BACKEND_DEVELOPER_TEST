package handler 

type RequestCreateEmploye struct {
	ID              uint     `json:"id"` 
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
} 

type RequestUpdateEmploye struct {
	ID              uint     `json:"id"` 
	FirstName       string   `json:"firstname"` 
	LastName        string   `json:"lastname"`
	HireDate        string   `json:"hiredate"`
	TerminationDate string   `json:"terminationdate"`
	Salary          string   `json:"salary"`
} 