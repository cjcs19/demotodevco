package app

import (
	"bitbucket.org/demotodevco/app/controllers"
	"github.com/julienschmidt/httprouter"
)

//GetNewRoutes Function to manage route
func GetNewRoutes() *httprouter.Router {

	r := httprouter.New()

	r.GET("/toroman/:num", controllers.ShowRoman)
	r.GET("/tonumeral/:roman", controllers.ShowNumeral)

	return r
}
