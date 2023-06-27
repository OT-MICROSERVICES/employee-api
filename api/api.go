package api

import (
	"employee-api/client"
	"employee-api/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// @Summary ReadEmployeesDesignation is a method to read all employee designation
// @Schemes http
// @Description Read all employee location data from database
// @Tags employee
// @Accept json
// @Produce json
// @Success 200 {object} model.Designation
// @Router /search/designation [get]
// ReadEmployeesDesignation is a method to read all employee location
func ReadEmployeesDesignation(c *gin.Context) {
	var employee model.Employee
	var designationResponse model.Designation
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	designation := make(map[string]int)
	data := scyllaClient.Query("SELECT designation FROM employee_info").Iter()
	for data.Scan(&employee.Designation) {
		_, exist := designation[employee.Designation]
		if exist {
			designation[employee.Designation] += 1
		} else {
			designation[employee.Designation] = 1
		}
	}
	jsonData, err := json.Marshal(designation)
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	json.Unmarshal(jsonData, &designationResponse)
	c.JSON(http.StatusOK, designationResponse)
}

// @Summary ReadEmployeesLocation is a method to read all employee location
// @Schemes http
// @Description Read all employee location data from database
// @Tags employee
// @Accept json
// @Produce json
// @Success 200 {object} model.Location
// @Router /search/location [get]
// ReadEmployeesLocation is a method to read all employee location
func ReadEmployeesLocation(c *gin.Context) {
	var employee model.Employee
	var locationResponse model.Location
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	location := make(map[string]int)
	data := scyllaClient.Query("SELECT office_location FROM employee_info").Iter()
	for data.Scan(&employee.OfficeLocation) {
		_, exist := location[employee.OfficeLocation]
		if exist {
			location[employee.OfficeLocation] += 1
		} else {
			location[employee.OfficeLocation] = 1
		}
	}
	jsonData, err := json.Marshal(location)
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	json.Unmarshal(jsonData, &locationResponse)
	c.JSON(http.StatusOK, locationResponse)
}

// @Summary ReadCompleteEmployeesData is a method to read all employee's information
// @Schemes http
// @Description Read all employee data from database
// @Tags employee
// @Accept json
// @Produce json
// @Success 200 {array} model.Employee
// @Router /search/all [get]
// ReadCompleteEmployeesData is a method to read all employee information
func ReadCompleteEmployeesData(c *gin.Context) {
	var employee model.Employee
	var response []model.Employee
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	data := scyllaClient.Query("SELECT id, name, designation, department, joining_date, address, office_location, status, email, annual_package, phone_number FROM employee_info").Iter()
	for data.Scan(&employee.ID, &employee.Name, &employee.Designation, &employee.Department, &employee.JoiningDate, &employee.Address, &employee.OfficeLocation, &employee.Status, &employee.EmailID, &employee.AnnualPackage, &employee.PhoneNumber) {
		response = append(response, employee)
	}
	c.JSON(http.StatusOK, response)
}

// @Summary ReadEmployeeData is a method to read employee information
// @Schemes http
// @Description Read data from database
// @Tags employee
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} model.Employee
// @Router /search [get]
// ReadEmployeeData is a method to read employee information
func ReadEmployeeData(c *gin.Context) {
	var response model.Employee
	mapData := map[string]interface{}{}
	id, keyExists := c.GetQuery("id")
	if !keyExists {
		logrus.Errorf("Query request of data without params")
		errorResponse(c, "Unable to perform search operation, query params not defined")
		return
	}
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in reading data from scylladb: %v", err)
		errorResponse(c, "Cannot read data from the system, request failure")
		return
	}
	data := scyllaClient.Query("SELECT * FROM employee_info where id = ?", id).Iter()
	for data.MapScan(mapData) {
		jsonData, err := json.Marshal(mapData)
		if err != nil {
			logrus.Errorf("Error in reading data from scylladb: %v", err)
			errorResponse(c, "Cannot read data from the system, request failure")
			return
		}
		json.Unmarshal(jsonData, &response)
	}
	c.JSON(http.StatusOK, response)
}

// @Summary CreateEmployeeData is a method to write employee information in database
// @Schemes http
// @Description Write data in database
// @Tags employee
// @Accept json
// @Produce json
// @Param employee body model.Employee true "Employee Data"
// @Success 200 {object} model.Employee
// @Router /create [post]
// CreateEmployeeData is a method to write employee information in database
func CreateEmployeeData(c *gin.Context) {
	var request model.Employee
	if err := c.BindJSON(&request); err != nil {
		logrus.Errorf("Error parsing the request body in JSON: %v", err)
		errorResponse(c, "Unable to Bind JSON in defined format, seems malformed")
		return
	}
	scyllaClient, err := client.CreateScyllaDBClient()
	if err != nil {
		logrus.Errorf("Error in writing data to scylladb: %v", err)
		errorResponse(c, "Cannot write data to the system, request failure")
		return
	}
	defer scyllaClient.Close()
	date, err := time.Parse("2006-01-02", request.JoiningDate)
	if err != nil {
		logrus.Errorf("Error in writing data to scylladb: %v", err)
		errorResponse(c, "Cannot write data to the system, request failure")
		return
	}
	queryString := "INSERT INTO employee_info(id, name, designation, department, joining_date, address, office_location, status, email, annual_package, phone_number) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	if err := scyllaClient.Query(queryString,
		request.ID, request.Name, request.Designation, request.Department, date, request.Address, request.OfficeLocation, request.Status, request.EmailID, request.AnnualPackage, request.PhoneNumber).Exec(); err != nil {
		logrus.Errorf("Error in writing data to scylladb: %v", err)
		errorResponse(c, "Cannot write data to the system, request failure")
		return
	}
	data := model.CustomMessage{
		Message: "Successfully created the data for the user",
	}
	logrus.Infof("Successfully created the employee record")
	c.JSON(http.StatusOK, data)
}

func errorResponse(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, model.CustomMessage{
		Message: err,
	})
}
