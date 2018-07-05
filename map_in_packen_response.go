package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bukalapak/packen/response"
	"github.com/go-sql-driver/mysql"
)

// handle response both success or failed
func HandleResponse(w http.ResponseWriter, response interface{}, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	switch status {
	case 200:
		WriteSuccess(w, response, status)
		break
	case 201:
		WriteSuccess(w, response, status)
		break
	default:
		WriteError(w, err, status)
		break
	}

}

// build error response and write it
func WriteError(w http.ResponseWriter, err string, status int) {
	errCust := response.CustomError{
		Message:  err,
		HTTPCode: status,
	}
	errs := []error{errCust}
	res := response.BuildError(errs)
	response.Write(w, res, status)
}

// build success response and write it
func WriteSuccess(w http.ResponseWriter, data interface{}, status int) {
	res := response.BuildSuccess(data, response.MetaInfo{HTTPStatus: status})
	response.Write(w, res, status)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// m := make(map[string]interface{})
	// m["id"] = "1001012"
	// m["a"] = "b"
	// c := make(map[string]interface{})
	// c["cliff"] = 123
	// c["luqman"] = 1
	// m["map"] = c

	a := 4
	alias := "lala"
	boo := false
	u := &DummyUser{
		ID:        3,
		Egg:       &a,
		Name:      "luqman",
		Alias:     &alias,
		IsMan:     true,
		Handsome:  &boo,
		UpdatedAt: mysql.NullTime{Time: time.Now()},
	}

	WriteSuccess(w, u, 200)
}

func MapInPackenResponse() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":1234", nil)
	fmt.Println("done")
}
