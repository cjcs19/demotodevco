package app

import (
	"github.com/demotodevco/app/controllers"
	"github.com/julienschmidt/httprouter"
)

//GetNewRoutes Function to manage route
func GetNewRoutes() *httprouter.Router {

	r := httprouter.New()

	r.GET("/toroman/:num", controllers.ShowRoman)
	r.GET("/tonumeral/:roman", controllers.ShowNumeral)

	return r
}
