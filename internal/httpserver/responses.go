package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// Error helps to return a graceful error
	Error struct {
		HttpStatus int    `json:"-"`
		Message    string `json:"message"`
		Code       string `json:"code"`
	}
)

// Error.Error Prettify the error for json format
func (e *Error) Error() string {
	fmtError, err := json.Marshal(e)
	if err != nil {
		return "Error on trying format http error"
	}
	return string(fmtError)
}

func (e *Error) String() string {
	return e.Error()
}
func (e *Error) ToHttpResponse(w http.ResponseWriter) error {
	return buildResponse(
		w,
		e,
		e.HttpStatus,
	)
}



// BuildOkResponse responses with status 200
func BuildOkResponse(w http.ResponseWriter, result fmt.Stringer) error {
	return buildResponse(w, result, http.StatusOK)
}

// BuildBadRequestResponse responses with status 400
func BuildBadRequestResponse(w http.ResponseWriter, result fmt.Stringer) error {
	return buildResponse(
		w,
		result,
		http.StatusBadRequest,
	)
}

func buildResponse(w http.ResponseWriter, result fmt.Stringer, statusCode int) error {
	w.WriteHeader(statusCode)
	_, err := fmt.Fprint(w, result.String())
	return err
}
