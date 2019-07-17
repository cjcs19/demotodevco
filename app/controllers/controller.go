package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cjcs19/demotodevco/app/conversor"

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

	results := map[string]interface{}{"De Romano A Numero-mejorado-Ingress-1": n}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(results); err != nil {
		return
	}

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

	//json.NewEncoder(w).Encode(results)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(results); err != nil {
		return
	}

}
