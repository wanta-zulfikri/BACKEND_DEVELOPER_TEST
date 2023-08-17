package handler

type ResponseGetAnnuals struct {
	ID     		uint	`json:"id"` 
	EmplID  	string 	`json:"emplid"`
	ReviewDate 	string 	`json:"reviewdate"`
}


type ResponseGetAnnual struct {
	ID 				uint     `json:"id"`
	EmplID  		string   `json:"emplid"`
	ReviewDate		string   `json:"reviewdate"`
} 

type AnnualResponse struct {
	Code       int         `json:"code"`
    Message    string      `json:"message"`
	Data       AnnualData  `json:"data"`
}

type AnnualData   struct {
	ID 				uint     `json:"id"`
	EmplID  		string   `json:"emplid"`
	ReviewDate		string   `json:"reviewdate"`
}  

type AnnualsResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"` 
	Data  []ResponseGetAnnuals  `json:"data"`
	Pagination  Pagination      `json:"pagination"`
} 

type Pagination struct {
	Page       int   `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}  

type ResponseUpdateAnnuals struct {
	ID     		uint	`json:"id"` 
	EmplID  	string 	`json:"emplid"`
	ReviewDate 	string 	`json:"reviewdate"`
}