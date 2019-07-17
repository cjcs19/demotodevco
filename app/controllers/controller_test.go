// handlers_test.go
package controllers

import (
	//"fmt"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGetRoman(t *testing.T) {
	req1, err := http.NewRequest("GET", "/toroman/100", nil)
	if err != nil {
		t.Fatal(err)
	}

	//fmt.Printf("Result RQ1: %+v\n", req1)

	rr1 := newRequestRecorder(req1, "GET", "/toroman/:num", ShowRoman)

	//fmt.Printf("Resultrr1: %v\n", rr1.Body.String())

	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	expected := `{"De Numero A Romano":"C"}`
	body := strings.TrimSpace(rr1.Body.String())

	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr1.Body.String(), expected)
	}
}

func TestGetNumeral(t *testing.T) {
	req1, err := http.NewRequest("GET", "/tonumeral/XL", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr1 := newRequestRecorder(req1, "GET", "/tonumeral/:roman", ShowNumeral)

	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	//fmt.Printf("Result RQ1: %+v\n", req1)

	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	expected := `{"De Romano A Numero-mejorado-Ingress-6":40}`
	body := strings.TrimSpace(rr1.Body.String())

	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr1.Body.String(), expected)
	}

}

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	router.ServeHTTP(rr, req)
	return rr
}
