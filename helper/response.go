package helper

type DataResponse struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
} 

func ResponseFormat(code int, message string, data interface{}) map[string]interface{}{
		res := make(map[string]interface{}) 
		res["code"] = code 
		if message != "" {
				res["message"] = message
		} else {
			res["message"] = "Successful Operation"
		} 
		if data != nil {
			 	res["data"] = data
		}
		return res
}