package handler

type RequestCreateAnnual struct {
	ID     		uint	`json:"id"` 
	EmplID  	string 	`json:"emplid"`
	ReviewDate 	string 	`json:"reviewdate"`
}


type RequestUpdateAnnual struct {
	ID 				uint     `json:"id"`
	EmplID  		string   `json:"emplid"`
	ReviewDate		string   `json:"reviewdate"`
}