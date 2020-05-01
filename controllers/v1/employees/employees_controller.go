package employeescontroller

import (
	"net/http"
	"time"

	dbs "../../../config/db"
	"../../../config/responses"
	"../../../libs/files"
	httplib "../../../libs/http"
	employeesmodel "../../../models/employees"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

var (
	env          = viper.GetString("env")
	dbName       = "vms"
	dbCollection = "visitors"
)

//RegisterEmployee controller

func UploadImage(res http.ResponseWriter, req *http.Request) {
	url := files.UploadFile("image", req)

	resp := responses.GeneralResponse{Success: true, Data: url, Message: "image uploaded"}

	httplib.Response(res, resp)
}

func RegisterEmployee(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	var data employeesmodel.Employees

	c.BindJSON(&data)

	data.ID = bson.NewObjectId()
	data.Date = time.Now()

	err := coll.Insert(data)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating employee"}
		httplib.Response(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: data, Message: "employee created"}
	httplib.Response(res, resp)
}

//GetEmployeeDetails controller
func GetEmployeeDetails(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	employeeEmail := c.Params("employeeEmail")

	var employee interface{}

	err := coll.Find(bson.M{"email": employeeEmail}).One(&employee)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error getting employee"}
		httplib.Response(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: employee, Message: "employee details"}
	httplib.Response(res, resp)
}

//UpdateEmployeeDetails controller
func UpdateEmployeeDetails(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}
	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	var updates bson.M

	c.BindJSON(updates)
	employeeID := c.Params("employeeEmail")

	err := coll.Update(bson.M{"_id": bson.ObjectIdHex(employeeID)}, bson.M{"$set": updates})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error updating employee"}
		httplib.Response(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: updates, Message: "employee updated"}
	httplib.Response(res, resp)
}

//DeleteEmployee controller
func DeleteEmployee(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	employeeEmail := c.Params("employeeEmail")

	err := coll.Remove(bson.M{"email": employeeEmail})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error deleting employee"}
		httplib.Response(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: employeeEmail, Message: "employee deleted"}
	httplib.Response(res, resp)
}
