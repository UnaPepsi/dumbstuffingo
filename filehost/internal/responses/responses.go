package responses

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
)
type Response interface {
	ErrorResponse | LoginResponse | FileUploadedResponse
}

type ErrorResponse struct {
	Message string `json:"message"`
	Ratelimit int `json:"ratelimit"`
}

func (e *ErrorResponse) Error() string{
	return e.Message
}

func SendResponse[T Response](r *T, w *http.ResponseWriter, statusCode int16){
	out, err := json.Marshal(*r)
	if err != nil {
		log.Fatalf("An error ocurred during json parsing: %v", err.Error())
	}
	(*w).WriteHeader(int(statusCode))
	fmt.Fprint(*w,out)

}
type LoginResponse struct {
	Token string `json:"token"`
}

type FileUploadedResponse struct {
	Id int `json:"id"`
}
