package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"bitbucket.org/demotodevco/app/conversor"

	"github.com/julienschmidt/httprouter"
)

func ShowNumeral(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//fmt.Printf("\nEntro a ShowNumeral\n")

	stringRoman := ps.ByName("roman")
	conversortoNumeral := conversor.NewRoman()
	n := conversortoNumeral.ToNumber(stringRoman)

	/*if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/

	results := map[string]interface{}{"De Romano A Numero": n}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(results)

}

func ShowRoman(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//fmt.Printf("\nEntro a ShowRoman %+v\n", ps)

	number, _ := strconv.Atoi(ps.ByName("num"))
	//fmt.Printf("\n Number a ShowRoman %v\n", number)

	conversortoRoman := conversor.NewRoman()

	n := conversortoRoman.ToRoman(number)
	//fmt.Printf("\nResultado Toroman %v\n", n)

	/*if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/

	results := map[string]interface{}{"De Numero A Romano": n}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(results)

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}