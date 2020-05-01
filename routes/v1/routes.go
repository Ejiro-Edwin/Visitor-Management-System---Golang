package routes

import (
	"net/http"

	"../../config/responses"
	employeecontoller "../../controllers/v1/employees"
	visitorscontoller "../../controllers/v1/visitors"
	visitscontroller "../../controllers/v1/visits"
	httplib "../../libs/http"
	mws "../../middlewares"
	"github.com/gorilla/mux"
)

//Router for all routes
func Router() *mux.Router {
	route := mux.NewRouter()

	//BASE ROUTE
	route.HandleFunc("/v1", func(res http.ResponseWriter, req *http.Request) {
		resp := responses.GeneralResponse{Success: true, Message: "vms  server running....", Data: "vsm SERVER v1.0"}
		httplib.Response(res, resp)
	})

	route.Use(mws.AccessLogToConsole)

	//************************
	// VISITES  ROUTES
	//************************
	visitsRoute := route.PathPrefix("/v1/visits").Subrouter()
	visitsRoute.HandleFunc("", visitscontroller.CreateVisits).Methods("POST")
	visitsRoute.HandleFunc("/{employeeEmail}", visitscontroller.GetEmployeeVisites).Methods("GET")
	visitsRoute.HandleFunc("/{visitorEmail}", visitscontroller.GetVisitorVisites).Methods("GET")
	visitsRoute.HandleFunc("/{visitorCode}", visitscontroller.Checkout).Methods("PUT")

	//************************
	// VISITORS  ROUTES
	//************************
	visitorsRoute := route.PathPrefix("/v1/visitors").Subrouter()
	visitorsRoute.HandleFunc("", visitorscontoller.RegisterVistor).Methods("POST")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.GetVisitorDetailsByEmail).Methods("GET")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.UpdateVistorDetials).Methods("PUT")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.UploadImage).Methods("POST")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.DeleteVistor).Methods("DELETE")

	//************************
	// EMPLOYEE  ROUTES
	//************************
	employeeRoute := route.PathPrefix("/v1/employee").Subrouter()
	employeeRoute.HandleFunc("", employeecontoller.RegisterEmployee).Methods("POST")
	employeeRoute.HandleFunc("/{employeeEmail}", employeecontoller.GetEmployeeDetails).Methods("GET")
	employeeRoute.HandleFunc("/{employeeEmail}", employeecontoller.UpdateEmployeeDetails).Methods("PUT")
	employeeRoute.HandleFunc("/{employeeEmail}", employeecontoller.UploadImage).Methods("POST")
	employeeRoute.HandleFunc("/{employeeEmail}", employeecontoller.DeleteEmployee).Methods("DELETE")

	return route
}
